package utils

import (
	"bytes"
	"os"
	"testing"
)

const testDB = "testdata/blockchain.db"

func TestDatabase(t *testing.T) {

	os.Remove(testDB)

	db, err := OpenDB(testDB)
	if err != nil {
		t.Fatal(err)
	}

	if db.IsBucketExist() != false {
		t.Errorf("Expected 'Bucket is not exist'")
	}
	err = db.CreateNewBucket()
	if err != nil {
		t.Error(err)
	}
	if db.IsBucketExist() != true {
		t.Errorf("Expected 'Bucket is exist'")
	}
	//test values
	serial := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	hash := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}

	err = db.AddNewBlock(hash, serial)
	if err != nil {
		t.Error(err)
	}

	newSerial := db.GetBlockByHash(hash)
	if !bytes.Equal(newSerial, serial) {
		t.Errorf("Expected %v, received %v", serial, newSerial)
	}
	newHash, err := db.GetLastHash()
	if err != nil {
		t.Error(err)
	}
	if newHash != hash {
		t.Errorf("Expected %v, received %v", hash, newHash)
	}

	os.Remove(testDB)

}
