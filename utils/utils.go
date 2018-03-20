package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"

	"golang.org/x/crypto/sha3"
	"log"
)

// UintToHex converts an int64 to a byte array
func UintToHex(num uint64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func EncodeToBytes(data interface{}) []byte {
	var buf bytes.Buffer

	gob.Register([32]byte{})
	gob.Register([8]byte{})

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}

func Hash(data []byte) [32]byte {

	hf := sha3.New256()
	hf.Write(data)

	h := hf.Sum(nil)

	var hashArray [32]byte
	copy(hashArray[:], h[:])
	return hashArray
}

func EncodeAndHash(data interface{}) [32]byte {
	return Hash(EncodeToBytes(data))
}
