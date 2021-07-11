package main

import (
	"crypto/sha256"
	"fmt"
)

//1.定义结构（区块头的字段比正常的少）
//>1.前区块哈希
//>2.当前区块哈希
//>3.数据

//2.创建区块
//3.生成哈希
//4.引入区块链
//5.添加区块
//6.重构代码

const genesisInfo = "genesis block is created"

// Block struct
type Block struct {
	PrevBlockHash []byte
	Hash          []byte
	Data          []byte //数据，目前使用字节流，v4开始使用交易代替
}

// 创建区块，对Block的每一个字段填充数据即可

// NewBlock func
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
	}

	block.SetHash()
	return block
}

//为了生成区块哈希，我们实现一个简单的函数，来计算hash值

// SetHash func
func (block *Block) SetHash() {
	var data []byte
	data = append(data, block.PrevBlockHash...)
	data = append(data, block.Data...)
	hash /*[32]byte*/ := sha256.Sum256(data)
	block.Hash = hash[:]
}

//创建区块链，使用Block数组模拟

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

func main() {
	// block := NewBlock(genesisInfo, []byte{0x0000000000000000})
	bc := NewBlockChain()
	bc.AddBlock("班主任来了，大家欢迎")
	for i, block := range bc.Blocks {
		fmt.Printf("+++++++++%d++++++++\n", i)
		fmt.Printf("PrevBlockHash:%x\n", block.PrevBlockHash)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("Data:%s\n", block.Data)
	}
}
