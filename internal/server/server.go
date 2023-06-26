package server

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/namelew/DHashTable/packages/hashtable"
)

type FileSystem struct {
	adress string
	id     uint64
	start  int
	end    int
	inodes hashtable.HashTable[string, string]
}

const SOURCEFILE = "./routing_table.in"

func removeBackSlash(s string) string {
	backslash := []string{"\n", "\a", "\b", "\r"}

	for i := range backslash {
		s = strings.ReplaceAll(s, backslash[i], "")
	}

	return s
}

func New(id uint64, adress string) *FileSystem {
	data, err := os.ReadFile(SOURCEFILE)

	if err != nil {
		log.Panic("Unable to create file system. Error on sourcefile read: ", err.Error())
	}

	lines := strings.Split(string(data), "\n")

	size, err := strconv.Atoi(removeBackSlash(lines[0]))

	if err != nil {
		log.Panic("Unable to create file system. Error on table size load: ", err.Error())
	}

	var start, end int = 0, 0

	for i := range lines {
		if i > 0 {
			cols := strings.Split(lines[i], " ")

			if len(cols) < 3 {
				continue
			}

			sid, err := strconv.ParseUint(removeBackSlash(cols[0]), 10, 64)

			if err != nil {
				log.Panic("Unable to create file system. Error on start and end load: ", err.Error())
			}

			if sid == id {
				start, err = strconv.Atoi(removeBackSlash(cols[1]))

				if err != nil {
					log.Panic("Unable to create file system. Error on table start load: ", err.Error())
				}

				end, err = strconv.Atoi(removeBackSlash(cols[2]))

				if err != nil {
					log.Panic("Unable to create file system. Error on table end load: ", err.Error())
				}
			}
		}
	}

	return &FileSystem{
		id:     id,
		adress: adress,
		start:  start,
		end:    end,
		inodes: hashtable.New[string, string](&hashtable.Open[string, string]{}, hashtable.Common{Size: size}),
	}
}
