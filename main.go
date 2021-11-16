package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	erc20 "demo/erc20"
)

func main() {
	client, err := ethclient.Dial("https://rpc-mainnet.maticvigil.com/")
	if err != nil {
		log.Fatal(err)
	}
	// define variable
	address := common.HexToAddress("0x6ae7dfc73e0dde2aa99ac063dcf7e8a63265108c")
	// account := common.HexToAddress("0x3f429e7ca142c81a1dba08f5264952503b28a933")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(21420146),
		ToBlock:   big.NewInt(21420146),
		Addresses: []common.Address{
			address,
		},
	}
	// get logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	// parse contract ABI
	contractAbi, err := abi.JSON(strings.NewReader(string(erc20.Erc20ABI)))
	fmt.Println("contractAbi", contractAbi)
	if err != nil {
		log.Fatal(err)
	}
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		fmt.Println("vLog", vLog.Topics[0])
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			// var transferEvent erc20.Erc20Transfer

			dataRaw, err := contractAbi.Unpack("Transfer", vLog.Data)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("transfer value", dataRaw[0])

			// 	transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			// 	transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			// 	fmt.Printf("From: %s\n", transferEvent.From.Hex())
			// 	fmt.Printf("To: %s\n", transferEvent.To.Hex())
			// 	fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

			// case logApprovalSigHash.Hex():
			// 	fmt.Printf("Log Name: Approval\n")

			// 	var approvalEvent LogApproval

			// 	err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}

			// 	approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			// 	approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			// 	fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			// 	fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			// 	fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}

}
