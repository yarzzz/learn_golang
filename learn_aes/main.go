package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// 加密
func encryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	fmt.Printf("encrypt block size: %d\n", block.BlockSize())
	// src = padding(src, block.BlockSize())
	blockMode := cipher.NewCTR(block, key[:16])
	blockMode.XORKeyStream(src, src)
	return src, nil
}

// 解密
func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	fmt.Printf("decrypt block size: %d\n", block.BlockSize())
	blockMode := cipher.NewCTR(block, key[:16])
	blockMode.XORKeyStream(src, src)
	// src = unpadding(src)
	return src, nil
}

func test1() {
	key := make([]byte, 32)
	n, err := rand.Read(key)
	if err != nil {
		log.Fatalln(err)
	}
	key, _ = hex.DecodeString("603deb1015ca71be2b73aef0857d77811f352c073b6108d72d9810a30914dff4")
	fmt.Printf("key len: %d\n", n)
	d := []byte("qVizLCeKIZIPjJQ1QhM0ciXwv23EZOHQGbUJOPsw6Fr8gBzKQxBezupFaC2LKWo5")
	fmt.Printf("加密前: %x,len: %d\n", d, len(d))
	x1, err := encryptAES(d, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("加密后: %x,len: %d\n", x1, len(x1))
	x2, err := decryptAES(x1, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("解密后: %x,len: %d\n", x2, len(x2))
}

func main() {
	test1()
}
