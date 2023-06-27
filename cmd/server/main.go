package main

import (
	"github.com/namelew/DHashTable/internal/server"
)

func main() {
	pid := server.New(0)
	pid.Build()
}
