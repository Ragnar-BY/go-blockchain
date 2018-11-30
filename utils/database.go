package utils

import (
	"log"

	"github.com/boltdb/bolt"
)

const blocksBucket = "blocks"

// Database is database with bolt.DB.
// bucket has two key->values: l->lastHash,blockhash->block.Serialize()
type Database struct {
	*bolt.DB
}

// OpenDB creates and opens bolt.DB
func OpenDB(dbFile string) (*Database, error) {

	db, err := bolt.Open(dbFile, 0600, nil)
	return &Database{db}, err
}

// IsBucketExist checks if bucket exist
func (db *Database) IsBucketExist() bool {
	var b *bolt.Bucket
	_ = db.View(func(tx *bolt.Tx) error {
		b = tx.Bucket([]byte(blocksBucket))
		return nil
	})
	return !(b == nil)
}

// CreateNewBucket creates new bucket
func (db *Database) CreateNewBucket() error {

	err := db.Update(func(tx *bolt.Tx) error {

		log.Println("No existing blockchain found. Creating a new one...")
		_, err := tx.CreateBucket([]byte(blocksBucket))
		return err
	})
	return err
}

// GetLastHash gets lash hash value from db.
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

// GetBlockByHash gets block by hash or return nil if not found
func (db *Database) GetBlockByHash(hash [32]byte) []byte {

	var block []byte
	_ = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		block = b.Get(hash[:])
		return nil
	})
	return block
}

// AddNewBlock adds new block to blockchain.
func (db *Database) AddNewBlock(hash [32]byte, serial []byte) error {

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		h := hash[:]

		// add block
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
