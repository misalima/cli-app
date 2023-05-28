package main

import (
	"fmt"
	"os"
	"log"
	"github.com/misalima/cli-app/app"
)

func main() {
	fmt.Println("Ponto de partida")

	app := app.Generate()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}