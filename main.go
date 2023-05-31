package main

import (
	"os"
	"log"
	"github.com/misalima/cli-app/app"
)

func main() {
	app := app.Generate()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}