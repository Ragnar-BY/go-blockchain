package main

import (
	"fmt"
	"time"

	"github.com/Ragnar-BY/go-blockchain/types"
)

func main() {

	fmt.Println("Start " + time.Now().Format("2006-01-02 15:04:05.99"))
	bc := types.NewBlockChain()
	defer bc.CloseDB()

	bc.AddBlock([]byte("First block"))
	bc.AddBlock([]byte("Second block"))
	bc.AddBlock([]byte("Third block"))

	fmt.Println("End " + time.Now().Format("2006-01-02 15:04:05.99"))

}
