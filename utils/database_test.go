package utils

import (
	"bytes"
	"io"
	"os"
	"testing"
)

const testNewDB = "testdata/blockchain.db"
const testExistDB = "testdata/testblockchain.db"

//test if there are not db file
func TestNewDatabase(t *testing.T) {

	os.Remove(testNewDB)

	db, err := OpenDB(testNewDB)
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

	os.Remove(testNewDB)
}

//test if there is db file with data
func TestExistingDatabase(t *testing.T) {

	from, err := os.Open(testExistDB)
	dbFile := testExistDB + ".copy"
	if err != nil {
		t.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(dbFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		t.Fatal(err)
	}

	db, err := OpenDB(dbFile)
	if err != nil {
		t.Fatal(err)
	}

	if db.IsBucketExist() != true {
		t.Errorf("Expected 'Bucket is exist'")
	}
	//test values( this data are in Bucket
	serial := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	hash := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}

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

	serial2 := []byte{8, 7, 6, 5, 4, 3, 2, 1}
	hash2 := [32]byte{8, 7, 6, 5, 4, 3, 2, 1, 8, 7, 6, 5, 4, 3, 2, 1, 8, 7, 6, 5, 4, 3, 2, 1, 8, 7, 6, 5, 4, 3, 2, 1}

	err = db.AddNewBlock(hash2, serial2)
	if err != nil {
		t.Error(err)
	}

	newSerial = db.GetBlockByHash(hash2)
	if !bytes.Equal(newSerial, serial2) {
		t.Errorf("Expected %v, received %v", serial2, newSerial)
	}
	newHash, err = db.GetLastHash()
	if err != nil {
		t.Error(err)
	}
	if newHash != hash2 {
		t.Errorf("Expected %v, received %v", hash2, newHash)
	}

	os.Remove(dbFile)
}
