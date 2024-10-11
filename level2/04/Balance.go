package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//client, err := ethclient.Dial("https://mainnet.infura.io/v3/<API-KEY>")
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/<your-infura-api-key>")

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x8Aa9f2cc55a5B9994b2B66c8Fb4b69E0ac6bcc6F")

	fmt.Println(account.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(account.Bytes())

	fmt.Println("we have a connection")

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 25893180161173005034
	_ = client           // we'll use this in the upcoming sections
}
