package main 

//targetBits defines how difficult mining is
//it is the value differs from 256, ensure this value is significant enough but not too big
//because the bigger the value the more difficult to find a proper hash
const targetBits = 24

type ProofOfWork struct {
    //ProofOfWork holds a pointer to a Block and a pointer to target
    block *Block
    target *big.Int

    //target is the requirements that hash needs to meet for instance “first 20 bits of a hash must be zeros”
    //a hash will be compared to the target in terms of the value 
    //that a hash converted into a big integer and this value is smaller than the target
    
    //target works as an upper boundry of a range, a hash is valid only when it is smaller than target
    //lowering target will result in fewer valid numbers which makes mining more difficult
}

func NewProofOfWork(b *Block) *ProofOfWork{
    //init a big integer with value of 1
    target := big.NewInt(1)

    //shift the bif integer by 256 - targetBits, 256 is the length of a SHA-256 hash in bits
    target.Lsh(target, uint(256-targetBits))

    pow := &ProofOfWork{b, target}

    return pow
}

func(pow *ProofOfWork) prepareData(nonce int) []byte{

    //merge block fields with target and nonce
    data := bytes.join(
                      [][]byte{ 
                               pow.block.PrevBlockHash,
                               pow.block.Data,
                               IntToHex(pow.block.Timestamp),
                               IntToHex(int64(targetBits)),
                               IntToHex(int64(nonce)),
                               },
                      []byte{},
                      )
    //nonce is the counter from Hashcash
    return data
}

//Run() performs a proof of work
func(pow *ProofOfWork) Run() (int, []byte){
    var hashInt big.Int
    var hash [32]byte
    nonce := 0
    fmt.Printf("mining the block containing \"%s\"\n", pow.block.Data)

    for nonce < MaxNonce {
        //prepare data
        data := pow.prepareData(nonce)

        //hash the data with SHA-256
        hash = sha256.Sum256(data)
        fmt.Printf("\r%x",hash)

        //convert the hash to big integer
        hashInt.SetBytes(hash[:])

        //compare the integer with the target
        if hashInt.Com(pow.target) == -1{
                break
            }else{
                nonce ++
            }
    }
    fmt.Print("\n\n")
    return nonce, hash[:]
}




























