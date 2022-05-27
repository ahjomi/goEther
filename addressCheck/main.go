package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Using simple regular expression to check if the ethereum address is valid
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}")

	fmt.Printf("is valid: %v\n", re.MatchString("0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"))

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// Checking if the address is a smart contract
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract)

	address = common.HexToAddress("0x22d491Bde2303f2f43325b2108D26f1eAbA1e32b")
	bytecode, err = client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract)

}
