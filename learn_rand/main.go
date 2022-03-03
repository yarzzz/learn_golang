package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func main() {
	datakey := make([]byte, 32)
	v, err := rand.Reader.Read(datakey)
	if err != nil {
		panic(err)
	}
	log.Println(v)
	log.Println(hex.EncodeToString(datakey))
}
