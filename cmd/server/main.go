package main

import (
	"log"

	"github.com/namelew/DHashTable/internal/server"
)

func main() {
	pid := server.New(0)
	log.Println(pid)
	pid.Build()
}
