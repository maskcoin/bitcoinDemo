package main

import "crypto/sha256"

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
