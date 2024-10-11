package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/224c8f29ab41435f92731adfebabe887")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f86d07830f42ef82520894efc3aa7241cd0d012f57950e7312e7b7bd0b97d1872386f26fc10000808401546d71a09e4dbe7f695211eef3e4cbe62bf1b680cf76a40ace36d473fd32b2521a71960ca071a6229791b2eb8fb0cc1161b0ea15835cded9b49f2da29f741b2e6417582723"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
