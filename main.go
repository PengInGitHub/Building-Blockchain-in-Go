package main 

import (
"fmt"

)

func main(){
    bc := NewBlockchain()
    bc.AddBlock("Send 1 BTC to Peng")
    bc.AddBlock("Send 2 BTC to Peng")

    for _,block := range bc.blocks{
        fmt.Printf("Prev. Hash: %x\n",block.PrevBlockHash)
        fmt.Printf("Data: %s\n",block.Data)
        fmt.Printf("Hash: %s\n",block.Hash)
        fmt.Println()
    }
}