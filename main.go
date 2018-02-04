package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func connect() {
	url := "https://eth.etherblue.org"
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println("Failed to dial, url: ", url, ", err: ", err)
		return
	}
	fmt.Println("Connected to eth.etherblue.org")

	address := common.StringToAddress("0xa8b8d9b1425ad962dce1b9606af606b6fd490037")

	client.BalanceAt(context.Background(), address, nil)
}

func main() {
	randBytes := make([]byte, 64)
	_, err := rand.Read(randBytes)
	if err != nil {
		panic("key generation: could not read from random source: " + err.Error())
	}
	reader := bytes.NewReader(randBytes)
	key, err := ecdsa.GenerateKey(crypto.S256(), reader)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}

	privateKey := fmt.Sprintf("%x", crypto.FromECDSA(key))
	addressRaw := crypto.PubkeyToAddress(key.PublicKey)
	address := addressRaw.String()

	fmt.Println("Private Key: ", privateKey)
	fmt.Println("Public Key: ", address)

	connect()
}
