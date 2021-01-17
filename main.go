package main

import (
	"findIp/app"
	"log"
	"os"
)

func main() {
	application := app.Build()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
