package server

import "github.com/namelew/DHashTable/packages/hashtable"

type FileSystem struct {
	hashTable hashtable.HashTable[string, string]
}
