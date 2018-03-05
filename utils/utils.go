package utils

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/crypto/sha3"
	"log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func Hash(data []byte) [32]byte {
	hf := sha3.New256()
	hf.Write(data)

	h := hf.Sum(nil)

	var hashArray [32]byte
	copy(hashArray[:], h[:])
	return hashArray
}
