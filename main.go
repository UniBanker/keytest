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

func balanceIsZero(client *ethclient.Client, address string) bool {
	addy := common.HexToAddress(address)

	bigIntBal, err := client.BalanceAt(context.Background(), addy, nil)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println("Balance: ", bigIntBal)
	return bigIntBal.Int64() == 0
}

func openGethClient() *ethclient.Client {
	url := "https://eth.etherblue.org"
	client, err := ethclient.Dial(url)
	if err != nil {
		fmt.Println("Failed to dial, url: ", url, ", err: ", err)
		return nil
	}
	fmt.Println("Connected to eth.etherblue.org")
	return client
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

	client := openGethClient()

	balZero := balanceIsZero(client, "0xA8b8d9B1425aD962dce1b9606AF606B6Fd490037")
	fmt.Println("Balance zero: ", balZero)
}
