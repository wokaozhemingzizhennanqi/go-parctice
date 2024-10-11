package main

import (
	"Practice/level2/store"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/224c8f29ab41435f92731adfebabe887")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x5807bc8FE545feccB880524B25740c172b5a2F1C")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
