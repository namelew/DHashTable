package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/namelew/DHashTable/packages/hashtable"
	"github.com/namelew/DHashTable/packages/messages"
)

type FileSystem struct {
	adress       string
	id           uint64
	start        int
	end          int
	lock         sync.Mutex
	neighborhood [][]int
	inodes       hashtable.HashTable[string, string]
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
	table := make([][]int, 0)

	for i := range lines {
		if i > 0 {
			cols := strings.Split(lines[i], " ")

			if len(cols) < 3 {
				continue
			}

			line := make([]int, 2)

			start, err = strconv.Atoi(removeBackSlash(cols[1]))

			if err != nil {
				log.Panic("Unable to create file system. Error on table start load: ", err.Error())
			}

			end, err = strconv.Atoi(removeBackSlash(cols[2]))

			if err != nil {
				log.Panic("Unable to create file system. Error on table end load: ", err.Error())
			}

			line[0] = start
			line[1] = end

			table = append(table, line)
		}
	}

	return &FileSystem{
		id:           id,
		adress:       adress,
		start:        table[id][0],
		end:          table[id][1],
		neighborhood: table,
		inodes:       hashtable.New[string, string](&hashtable.Open[string, string]{}, hashtable.Common{Size: size}),
	}
}

func (fs *FileSystem) insert(m *messages.Message) messages.Message {
	fs.lock.Lock()
	defer fs.lock.Unlock()
	return messages.Message{}
}

func (fs *FileSystem) query(m *messages.Message) messages.Message {
	fs.lock.Lock()
	defer fs.lock.Unlock()
	return messages.Message{}
}

func (fs *FileSystem) remove(m *messages.Message) messages.Message {
	fs.lock.Lock()
	defer fs.lock.Unlock()
	return messages.Message{}
}

func (fs *FileSystem) handlerRequests() {
	listener, err := net.Listen("tcp", fs.adress)

	if err != nil {
		log.Panic("Unable to create request handler: ", err.Error())
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Enable to create connection: ", err.Error())
			continue
		}

		go func(c net.Conn) {
			var request, response messages.Message

			if err := request.Receive(c); err != nil {
				log.Println("Unable to receive message from ", c.RemoteAddr().String(), ":", err.Error())
				return
			}

			switch request.Action {
			case messages.INSERT:
				response = fs.insert(&request)
			case messages.QUERY:
				response = fs.query(&request)
			case messages.REMOVE:
				response = fs.remove(&request)
			}

			if err := response.Send(c); err != nil {
				log.Println("Unable to send response to ", c.RemoteAddr().String(), ":", err.Error())
			}
		}(conn)
	}
}

func (fs *FileSystem) Build() {
	log.Println("Init request handler...")
	go fs.handlerRequests()

	log.Printf("File System Node %d started\n", fs.id)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
