package main

import (
	"log"

	"github.com/Ragnar-BY/go-blockchain/pkg/blockchain"
	"github.com/Ragnar-BY/go-blockchain/pkg/blockchain/database/bolt"
	"github.com/Ragnar-BY/go-blockchain/pkg/cli"
)

func main() {

	dbFile := "blockchain.db"
	db, err := bolt.OpenDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	bc, err := blockchain.NewBlockChain(db)
	if err != nil {
		log.Fatal(err)
	}
	c := cli.CLI{Blockchain: bc}
	err = c.Run()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
