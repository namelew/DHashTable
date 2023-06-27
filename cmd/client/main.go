package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/namelew/DHashTable/packages/messages"
)

func sanitaze(s string) string {
	trash := []string{"\n", "\b", "\r", "\t"}

	for i := range trash {
		s = strings.ReplaceAll(s, trash[i], "")
	}

	return s
}

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		p, err := r.ReadSlice('\n')

		if err != nil {
			log.Println(err.Error())
			continue
		}

		input := strings.Split(string(p), " ")

		if len(input) < 4 {
			continue
		}

		var m messages.Message

		a, err := strconv.Atoi(input[0])

		if err != nil {
			log.Println(err.Error())
			continue
		}

		m.Action = messages.Action(a)
		m.Key = input[2]
		m.Name = sanitaze(strings.Join(input[3:], " "))

		log.Println(m)
	}
}
