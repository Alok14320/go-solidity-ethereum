package main

// "github.com/ubiq/go-ubiq/ethclient"

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-solidity-ethereum/contract"
)

const (
	privateKey      = "e60c2586cae5c890f873b5b91b59b84c418d95ab5e12958d5dfd3116fba45689"     // replace your wallet private key
	contractAddress = "0x7dcde58338834b72759d33d0a9f8efd72912aaae"                           //replace with your contract address
	infuraurl       = "https://polygon-mumbai.infura.io/v3/6b3e26f3baa64014a909c1a86ab09877" // replace with your url
)

func main() {

	client, err := ethclient.Dial(infuraurl)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(1000000000) // 1 gwei

	// Connect to the deployed contract
	contractAddress := common.HexToAddress(contractAddress)
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tokenID := big.NewInt(1) // Replace with the desired token ID
	tx, err := instance.MintMTK(auth, tokenID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("log tx info", tx.Hash(), tx.ChainId().String())

}
