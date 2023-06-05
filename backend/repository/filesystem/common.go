package filesystem

import (
	"fmt"
	"io"
	"os"
	"time"
)

func backup(source string) error {
	// Skip if the file does not exist
	if _, err := os.Stat(source); err != nil {
		return nil
	}

	// Open source for reading
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	// Create new destination file
	version := time.Now().Format("20060102T150405")
	backupFilename := fmt.Sprintf("%s.%s", source, version)

	// if destination file exists generate a new one
	if _, err := os.Stat(backupFilename); err == nil {
		// filename failed, try to find a new one
		for i := 1; i < 11; i++ {
			fileName := fmt.Sprintf("%s-%d", backupFilename, i)
			_, err = os.Stat(fileName)
			if err != nil {
				backupFilename = fileName
				break
			}
		}
	}

	des, err := os.Create(backupFilename)
	if err != nil {
		return err
	}
	defer des.Close()

	// Back it up
	_, err = io.Copy(des, src)
	if err != nil {
		return err
	}

	return des.Sync()
}
