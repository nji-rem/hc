package main

import (
	"fmt"
	"log"
)

func main() {
	addr := fmt.Sprintf("tcp://:%d", 30001)

	application := InitializeApp()

	log.Fatalln(application.Bootstrap(addr))
}
