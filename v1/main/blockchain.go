package main

// BlockChain struct
type BlockChain struct {
	Blocks []*Block
}

// 实现创建区块链的方法

// NewBlockChain func
func NewBlockChain() *BlockChain {
	// 在创建的时候添加一个区块，创世块
	genesisBlock := NewBlock(genesisInfo, []byte{0x0000000000000000})
	bc := &BlockChain{
		Blocks: []*Block{genesisBlock},
	}
	return bc
}

// AddBlock func
func (bc *BlockChain) AddBlock(data string) {
	//1.创建一个区块
	prevBlockHash := bc.Blocks[len(bc.Blocks)-1].Hash
	block := NewBlock(data, prevBlockHash)
	//2.添加到bc.Blocks数组中
	bc.Blocks = append(bc.Blocks, block)
}
