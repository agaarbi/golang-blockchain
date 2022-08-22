package blockchain

// creates array of Block Addresses
type BlockChain struct {
	Blocks []*Block
}

// timestamp, block height etc are added when system becomes complicated
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Block is created on address
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// method of BlockChain struct
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// returns pointer to the Block created by CreateBlock
// using Genesis as the data and empty previous hash value
// Genesis means "the origin or mode of formation of something"
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// returns pointer addresses to the BlockChain struct
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
