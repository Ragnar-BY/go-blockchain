package utils

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testNewDB = "testdata/blockchain.db"
const testExistDB = "testdata/testblockchain.db"

func TestNewDatabase(t *testing.T) {

	db, err := OpenDB(testNewDB)
	require.NoError(t, err, "OpenDB error")
	defer func() {
		err = os.Remove(testNewDB)
		require.NoError(t, err, "Remove error")
	}()
	assert.False(t, db.IsBucketExist(), "bucket should not exist")

	err = db.CreateNewBucket()
	require.NoError(t, err, "CreateNewBucket error")
	assert.True(t, db.IsBucketExist(), "bucket should  exist")

	// test values
	serial := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	hash := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}

	err = db.AddNewBlock(hash, serial)
	require.NoError(t, err, "AddNewBlock error")

	newSerial := db.GetBlockByHash(hash)
	assert.Equal(t, serial, newSerial)

	newHash, err := db.GetLastHash()
	require.NoError(t, err, "GetLastHash error")
	require.Equal(t, hash, newHash)

}

func TestExistingDatabase(t *testing.T) {

	from, err := ioutil.ReadFile(testExistDB)
	require.NoError(t, err, "ReadFile error")
	dbFile := testExistDB + ".copy"

	err = ioutil.WriteFile(dbFile, from, 0644)
	require.NoError(t, err, "WriteFile error")

	db, err := OpenDB(dbFile)
	require.NoError(t, err, "OpenDB error")
	defer func() {
		err = os.Remove(dbFile)
		require.NoError(t, err, "Remove error")
	}()

	assert.True(t, db.IsBucketExist(), "bucket should  exist")

	// test values( this data are in Bucket
	serial := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	hash := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}

	newSerial := db.GetBlockByHash(hash)
	assert.Equal(t, serial, newSerial)

	newHash, err := db.GetLastHash()
	require.NoError(t, err, "GetLastHash error")
	require.Equal(t, hash, newHash)

	serial2 := []byte{8, 7, 6, 5, 4, 3, 2, 1}
	hash2 := [32]byte{8, 7, 6, 5, 4, 3, 2, 1, 8, 7, 6, 5, 4, 3, 2, 1, 8, 7, 6, 5, 4, 3, 2, 1, 8, 7, 6, 5, 4, 3, 2, 1}

	err = db.AddNewBlock(hash2, serial2)
	require.NoError(t, err, "AddNewBlock error")

	newSerial = db.GetBlockByHash(hash2)
	assert.Equal(t, serial2, newSerial)

	newHash, err = db.GetLastHash()
	require.NoError(t, err, "GetLastHash error")
	require.Equal(t, hash2, newHash)
}
