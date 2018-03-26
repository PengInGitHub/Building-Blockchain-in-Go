package main 

import(
    "log"
    "github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct{
    tip []byte
    db *bolt.DB
}

//An iterator will be created each time we iterate over blocks in a blockchain 
//and it’ll store the block hash of the current iteration and a connection to a DB
type BlockchainIterator struct{
    currentHash []byte
    db *bolt.DB
}


//AddBlock() saves provided data as block in the blockchain
func(bc *Blockchain) AddBlock(data string){

    var lastHash []byte

    //obtain last block hash from the DB to mine a new block hash
    err := bc.db.View(func(tx *bolt.Tx) error{
        b := tx.Bucket([]byte(blocksBucket))
        lastHash = b.Get([]byte("l"))

        return nil
    })

    if err != nil {
        log.Panic(err)
    }

    //after mining a new block, its serialized representation is saved into the db
    //and update the l key that now stores the new block's hash

    newBlock := NewBlock{data, lastHash}    

    err = bc.db.Update(func(tx *bolt.Tx) error{
        b := tx.Bucket([]byte(blocksBucket))
        err := b.Put(newBlock.Hash, newBlock.Serialize())
        if err != nil {
            log.Panic(err)
        }

        err = b.Put([]byte("l"), newBlock.Hash)
        if err != nil {
            log.Panic(err)
        }

        bc.tip = newBlock.Hash

        return nil
    })
}

//Iterator() 
func(bc *Blockchain) Iterator *BlockchainIterator{
    bci := &BlockchainIterator{bc.tip, bc.db}
    return bci
}

//Next() returns the next block from the tip, tip is an identifier of a blockchain 
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

//NewBlockchain() creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain{
    var tip []byte

    //open a BoltDB file
    db, err := bolt.Open(dbFile, 0600, nil)
    if err != nil {
        log.Panic(err)
    }

    //db.Update open a read-write transcation for putting genesis Block into the DB
    err = db.Update(func(tx *bolt.Tx) error{
        
        //obtain bucket stores the blocks
        b := tx.Bucket([]byte(blocksBucket))

        //check if there is a blockchain stored in it        
        if b = nil{

            //there is no existing blockchain
            fmt.Println("No existing blockchain found. Create a new one...")

            //create the genesis block
            genesis := NewGenesisBlock()

            b, err := tx.CreateBucket([]byte(blocksBucket))
            if err != nil{
                log.Panic(err)
            }

            //store the genesis block in the DB
            err = b.Put(genesis.Hash, genesis.Serialize())
            if err != nil{
                log.Panic(err)
            }

            //save the genesis block's hash as the last block's hash         
            err = b.Put([]byte("l"), genesis.Hash)
            if err != nil{
                log.Panic(err)
            }

            tip = genesis.Hash
        }else{
            //blockchain instance is not nil, there is a blockchain
            //create a new blockchain instance to the last block hash stored in the DB
            tip = b.Get([]byte("l")) 
        }

        return nil
    })

    if err != nil{
        log.Panic(err)
    }

    bc := Blockchain{tip, db}
    return *bc
}























