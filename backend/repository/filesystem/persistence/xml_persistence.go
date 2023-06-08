package persistence

import (
	"cryptotracker/entity"
	"encoding/xml"
	"fmt"
	"os"
)

type root struct {
	Wallets []entity.Wallet `xml:"wallet"`
}

type xmlPersistence struct {
	filePath string
}

func NewXmlPersistence(filePath string) FsPersistence {
	return &xmlPersistence{filePath: filePath}
}

func (x *xmlPersistence) Save(wallets []entity.Wallet) error {
	// backup previous file
	err := backup(x.filePath)
	if err != nil {
		return fmt.Errorf("error backing up: %v", err)
	}

	// save new file
	xmlNode := root{
		Wallets: wallets,
	}

	fileData, err := xml.Marshal(xmlNode)

	if err != nil {
		return err
	}

	return os.WriteFile(x.filePath, fileData, 0644)
}

func (x *xmlPersistence) Load() ([]entity.Wallet, error) {
	// Skip if the file does not exist
	if _, err := os.Stat(x.filePath); err != nil {
		return nil, nil
	}

	contents, err := os.ReadFile(x.filePath)
	if err != nil {
		return nil, err
	}

	var node root
	err = xml.Unmarshal(contents, &node)
	if err != nil {
		return nil, err
	}

	return node.Wallets, nil
}
