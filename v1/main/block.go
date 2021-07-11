package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

const genesisInfo = "genesis block is created"

// Block struct
type Block struct {
	Version       uint64 //区块版本号
	PrevBlockHash []byte
	MerkleRoot    []byte
	TimeStamp     uint64
	Difficulty    uint64
	Nonce         uint64
	Data          []byte
	Hash          []byte
}

// 创建区块，对Block的每一个字段填充数据即可

// NewBlock func
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Version:       0,
		PrevBlockHash: prevBlockHash,
		TimeStamp:     uint64(time.Now().Unix()),
		Difficulty:    10, //随便写的，v2再调整
		Nonce:         10,
		Data:          []byte(data),
	}

	block.SetHash()
	return block
}

//为了生成区块哈希，我们实现一个简单的函数，来计算hash值

// SetHash func
func (block *Block) SetHash() {
	tmp := [][]byte{
		uintToByte(block.Version),
		block.PrevBlockHash,
		block.MerkleRoot,
		uintToByte(block.TimeStamp),
		uintToByte(block.Difficulty),
		uintToByte(block.Nonce),
		block.Data,
	}

	data := bytes.Join(tmp, nil)

	hash /*[32]byte*/ := sha256.Sum256(data)
	block.Hash = hash[:]
}
