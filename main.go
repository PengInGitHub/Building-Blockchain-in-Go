package main 

import (
"fmt"
"strconv"
)

func main(){
    bc := NewBlockchain()
    bc.AddBlock("Send 1 BTC to Peng")
    bc.AddBlock("Send 6 more BTC to Peng")

    for _,block := range bc.blocks{
        fmt.Printf("Prev. Hash: %x\n",block.PrevBlockHash)
        fmt.Printf("Data: %s\n",block.Data)
        fmt.Printf("Hash: %s\n",block.Hash)
        pow := NewProofOfWork(block)
        fmt.Printf("Pow: %s\n",strconv.FormatBool(pow.Validate()))
        fmt.Println()
    }
}