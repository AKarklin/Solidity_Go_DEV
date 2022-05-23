package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/AKarklin/Solidity_Go_DEV/projects/Faucet/api" // this would be your generated smart contract bindings
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

	// Адрес запрашивающего кран эккаунта (3-й эккауте в ганашах)
	//privateKey, err := crypto.HexToECDSA("e4c1fd41b4c4ed68cee805e5538a3916ef07744f78ef9c4c8b0309d76ea78e5e")
	// Адрес запрашивающего кран эккаунта (4-й эккауте в ганашах)
	privateKey, err := crypto.HexToECDSA("0ac5bc6ff8eb1d8a95cb2dcf3b27a1027f6c78bd435bdae186de7d64077e97ff")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	//auth.GasPrice = big.NewInt(1000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	auth.GasPrice = gasPrice

	// Адрес контракта
	contr, err := api.NewApi(common.HexToAddress("0x1B21105e9AF53c71A6c8bc25144F0b83052C1Bf5"), client)
	if err != nil {
		panic(err)
	}

	// Запрашиваем кран
	kran := new(big.Int)
	kran.SetString("100000000000000000", 10)
	tx, err := contr.Withdraw(auth, kran)
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.Hash())

}
