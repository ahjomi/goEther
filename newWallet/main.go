package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {

	// GenerateKey generates a random private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Convert it to bytes with crypto/ecdsa package and using the FromECDSA method
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// Convert it to a hexadecimal string using hexutil package which provides the Encode method
	// which takes a byte slice. Then we strip off the 0x after it's hex encoded
	fmt.Println(hexutil.Encode(privateKeyBytes[2:]))

	// Public key is derived from the private key
	publicKey := privateKey.Public()

	//Converting to hex
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes[4:]))

	// Returns the public address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	// Take the last 40 characters (20 bytes) and prefix it with 0x
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
