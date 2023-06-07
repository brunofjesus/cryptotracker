package main

import (
	"cryptotracker/repository/filesystem"
	"cryptotracker/rest/router"
	"cryptotracker/service"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	trackerRepository, err := filesystem.NewRepository("file.xml")
	if err != nil {
		panic(err)
	}

	svc := service.NewWalletService(trackerRepository)

	err = router.Start(svc)
	if err != nil {
		panic(err)
	}
}
