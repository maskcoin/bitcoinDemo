package main

import "fmt"

//1.定义结构（区块头的字段比正常的少）
//>1.前区块哈希
//>2.当前区块哈希
//>3.数据

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
	return block
}

func main() {
	block := NewBlock(genesisInfo, []byte{0x0000000000000000})
	fmt.Printf("PrevBlockHash:%x\n", block.PrevBlockHash)
	fmt.Printf("Hash:%x\n", block.Hash)
	fmt.Printf("Data:%s\n", block.Data)
}
