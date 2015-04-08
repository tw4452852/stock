package main

import (
	"log"

	"github.com/tw4452852/stock/server"
)

func main() {
	if e := server.StartSever(); e != nil {
		log.Printf("main: %s\n", e)
	}
}
