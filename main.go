package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"demo/jpyc"
)

func main() {
	client, err := ethclient.Dial("https://rpc-mainnet.maticvigil.com/")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x6ae7dfc73e0dde2aa99ac063dcf7e8a63265108c")
	// account := common.HexToAddress("0x3f429e7ca142c81a1dba08f5264952503b28a933")
	instance, err := jpyc.NewJpycCaller(address, client)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("instance", bal)
}
