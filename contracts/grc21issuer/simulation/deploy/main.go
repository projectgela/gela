package main

import (
	"context"
	"fmt"
	"github.com/projectgela/gela/accounts/abi/bind"
	"github.com/projectgela/gela/common"
	"github.com/projectgela/gela/contracts/grc21issuer"
	"github.com/projectgela/gela/contracts/grc21issuer/simulation"
	"github.com/projectgela/gela/ethclient"
	"log"
	"math/big"
	"time"
)

func main() {
	fmt.Println("========================")
	fmt.Println("mainAddr", simulation.MainAddr.Hex())
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)

	// init grc21 issuer
	auth := bind.NewKeyedTransactor(simulation.MainKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = big.NewInt(210000000000000)
	grc21IssuerAddr, grc21Issuer, err := grc21issuer.DeployGRC21Issuer(auth, client, simulation.MinApply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main address", simulation.MainAddr.Hex(), "nonce", nonce)
	fmt.Println("===> grc21 issuer address", grc21IssuerAddr.Hex())

	auth.Nonce = big.NewInt(int64(nonce + 1))

	// init trc20
	grc21TokenAddr, grc21Token, err := grc21issuer.DeployGRC21(auth, client, "TEST", "GEL", 18, simulation.Cap, simulation.Fee)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===>  grc21 token address", grc21TokenAddr.Hex(), "cap", simulation.Cap)

	fmt.Println("wait 10s to execute init smart contract")
	time.Sleep(10 * time.Second)

	grc21Issuer.TransactOpts.Nonce = big.NewInt(int64(nonce + 2))
	grc21Issuer.TransactOpts.Value = simulation.MinApply
	grc21Issuer.TransactOpts.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	grc21Token.TransactOpts.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	grc21Token.TransactOpts.GasLimit = uint64(4000000)
	auth.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	// get balance init grc21 smart contract
	balance, err := grc21Token.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(simulation.Cap) != 0 {
		log.Fatal(err, "\tget\t", balance, "\twant\t", simulation.Cap)
	}
	fmt.Println("balance", balance, "mainAddr", simulation.MainAddr.Hex())

	// add trc20 list token grc21 issuer
	_, err = grc21Issuer.Apply(grc21TokenAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("wait 10s to add token to list issuer")
	time.Sleep(10 * time.Second)

	//check grc21 SMC balance
	balance, err = client.BalanceAt(context.Background(), grc21IssuerAddr, nil)
	if err != nil || balance.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance  in grc21Issuer SMC: ", err, "got", balance, "wanted", simulation.MinApply)
	}

	//check balance fee
	balanceIssuerFee, err := grc21Issuer.GetTokenCapacity(grc21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", simulation.MinApply)
	}
}
