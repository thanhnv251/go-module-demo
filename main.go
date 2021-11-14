package main

import (
	"demo/environment"
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
)

func main() {
	var environment environment.Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("environment", environment.Post)
	// tx := "0xbcced66ead17a300840faf31d5f677c9eaaa182b962eb1264d241bdd0f27059e"
	// fmt.Println("Hello")
	// fmt.Println(math.Add(1, 2))
	// blockchain.GetTxDetail(tx)
}
