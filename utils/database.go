package utils

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

//bucket has two key->values: l->lastHash,blockhash->block.Serialize()
type Database struct {
	*bolt.DB
}

func OpenDB() (*Database, error) {

	db, err := bolt.Open(dbFile, 0600, nil)
	return &Database{db}, err
}

//check if bucket exist
func (db *Database) IsBucketExist() bool {
	var b *bolt.Bucket
	db.View(func(tx *bolt.Tx) error {
		b = tx.Bucket([]byte(blocksBucket))
		return nil
	})
	if b == nil {
		return false
	}
	return true
}

func (db *Database) CreateNewBucket() error {

	err := db.Update(func(tx *bolt.Tx) error {

		log.Println("No existing blockchain found. Creating a new one...")
		_, err := tx.CreateBucket([]byte(blocksBucket))
		return err
	})
	return err
}

//get 'l' value
func (db *Database) GetLastHash() ([32]byte, error) {
	var lastHash [32]byte
	err := db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blocksBucket))
		copy(lastHash[:], b.Get([]byte("l")))
		return nil
	})
	if err != nil {
		return [32]byte{}, err
	}
	return lastHash, nil
}

func (db *Database) AddNewBlock(hash [32]byte, serial []byte) error {

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		h := hash[:]

		//add block
		err := b.Put(h, serial)
		if err != nil {
			return err
		}
		// change last hash
		err = b.Put([]byte("l"), h)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
