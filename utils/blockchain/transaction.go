package blockchain

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTxDetail(tx string) {
	txHash := common.HexToHash(tx)
	client, err := ethclient.Dial("https://rpc-mainnet.maticvigil.com/")
	if err != nil {
		log.Fatal(err)
	}
	txReceipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("txReceipt", txReceipt.Status)
}
