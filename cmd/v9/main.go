package main

import (
	"fmt"
	"log"
)

func main() {
	addr := fmt.Sprintf("tcp://:%d", 30001)

	application, err := InitializeApp()
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(application.Bootstrap(addr))
}
