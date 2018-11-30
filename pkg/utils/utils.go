package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"

	"golang.org/x/crypto/sha3"
)

// UintToHex converts an int64 to a byte array.
func UintToHex(num uint64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// EncodeToBytes encode data to bytes.
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

// Hash gets sha3 hash.
func Hash(data []byte) ([32]byte, error) {

	hf := sha3.New256()
	_, err := hf.Write(data)
	if err != nil {
		return [32]byte{}, err
	}

	h := hf.Sum(nil)

	var hashArray [32]byte
	copy(hashArray[:], h[:])
	return hashArray, nil
}

// EncodeAndHash encodes and hashs.
func EncodeAndHash(data interface{}) ([32]byte, error) {
	return Hash(EncodeToBytes(data))
}
