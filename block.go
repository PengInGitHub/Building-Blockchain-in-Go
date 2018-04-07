package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

//Serizalize() serizalizes the block
func (b *Block) Serialize() []byte {

	//declare a Buffer that will store the serizalized data
	var result bytes.Buffer

	//init an encoder that encodes the block
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(&b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

//DeserializeBlock() deserializes the block, receive a byte array as input and returns a Block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
