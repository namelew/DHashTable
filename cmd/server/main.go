package main

import (
	"log"

	"github.com/namelew/DHashTable/internal/server"
)

func main() {
	pid := server.New(1)
	log.Println(pid)
	pid.Build()
}
