package main

import (
	"fmt"
	"go-http-server/internal/pkg/app"
	"log"
)

func main() {
	a, err := app.New()
	if err != nil {
		fmt.Println(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}

}
