package main

import (
	"fmt"
	"go-blockchain/types"
)

func main() {

	bc := types.NewBlockChain()

	bc.AddBlock([]byte("First block"))
	bc.AddBlock([]byte("Second block"))

	for _, block := range bc.Blocks() {
		fmt.Printf("%v \n", block.ToString())

	}
}
