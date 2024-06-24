package main

import (
	"fmt"
	"hc/internal/app"
	"log"
)

func main() {
	addr := fmt.Sprintf("tcp://:%d", 3000)

	application, err := app.InitializeApp(addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(application.Start())
}
