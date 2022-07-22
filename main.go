package main

import (
	s "blockchain/source"
	"fmt"
	"time"
)

func NewBlock(data string, prevHashBlock []byte) *s.Block {
	block := &s.Block{time.Now().UTC().Unix(), []byte(data), prevHashBlock, []byte{}}
	block.SetHash()
	return block
}

func main() {
	fmt.Println([]byte("hello world"))
}
