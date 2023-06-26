package main

import (
	"github.com/namelew/DHashTable/internal/server"
)

func main() {
	pid := server.New(1, "localhost:30001")
	pid.Build()
}
