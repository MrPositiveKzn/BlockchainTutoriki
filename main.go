package main

import (
	"fmt"
	"strconv"
	)
func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 0.4 BTC to Tosya")
	bc.AddBlock("Send 3 more BTC to Tosya")
	bc.AddBlock("Send 0.3 more BTC to Tosya")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
