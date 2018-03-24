package main 

import (
    "bytes"
    "crypto/sha256"
    "strconv"
    "time"
)

type Block struct{
    Timestamp int64
    Data byte[]
    PreBlockHash byte[]
    Hash byte[]
}

func NewBlock(data string, preBlockHash []byte) *Block{
    block := &Block{time.Now().Unix(), []byte(data),
        preBlockHash, []byte{}}
    block.SetHash()
    return block
}

func(b *Block) SetHash(){
    timestamp := []byte{strcov.FormatInt(b.Timestamp,10)}
    headers := bytes.Join([][]byte{b.PreBlockhash, b.Data, timestamp}, []byte{})
    hash := sha256.Sum256(headers)
    b.Hash= hash[:]
}

func NewGenesisBlock() *Block{
    return NewBlock("Genesis Block",[]byte{})
}

