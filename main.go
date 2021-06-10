package main

import (
	"address-searcher/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starter Point")
	app := app.Generate()

	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
		panic("Error while runing the application")
	}
}
