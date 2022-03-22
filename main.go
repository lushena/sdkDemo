package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main1() {
	fmt.Println("simple chain sdk demo")
	//创建keystore
	ks := keystore.NewPlaintextKeyStore("./keystore")
	//创建账户
	acct, err := ks.NewAccount("123")
	if err != nil {
		fmt.Println("failed to create newaccount:", err)
	}
	fmt.Println("acct:", acct.Address.Hex())
}
func main2() {
	fmt.Println("simple chain sdk demo")
	//创建keystore
	ks := keystore.NewKeyStore("./keystore", keystore.LightScryptN, keystore.LightScryptP)
	//创建账户
	acct, err := ks.NewAccount("123456")
	if err != nil {
		fmt.Println("failed to create newaccount:", err)
	}
	fmt.Println("acct:", acct.Address.Hex())
}

const keyinfo = `{"address":"6f50182271a93ec15de223647112b34693e4c595","crypto":{"cipher":"aes-128-ctr","ciphertext":"7c0cfea0d2b8b828fb1b14cf459060c2131dcf4fe09a84afa4335d3b74bfe77d","cipherparams":{"iv":"33326154fad04bcdbb3f908d72845bb5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"dc3d2f140ff2186374dad5d5d473a25a8cef6a77190f878409efac1431ff2fd4"},"mac":"0cda5d725eaeb6773a0cefceb0e64d4ac9ee5fcb0798bd0505def31e2060cb06"},"id":"972c3019-6bb7-4904-9475-f774c380cf32","version":3}`

func main3() {
	//1、创建客户端
	cli, err := ethclient.Dial("http://localhost:8541")
	if err != nil {
		log.Panic("failed to dial", err)
	}

	//2.身份准备
	keyin := strings.NewReader(keyinfo)
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Panic("failed to get chainID", err)
	}

	// auth, err := bind.NewTransactor(keyin, "123456")
	auth, err := bind.NewTransactorWithChainID(keyin, "123456", chainID)
	if err != nil {
		log.Panic("auth create failed:", err)
	}
	fmt.Println("chainID: ", chainID)

	//3、部署合约
	//auth *bind.TransactOpts, backend bind.ContractBackend, _msg string
	addr, tx, instance, err := DeployHello(auth, cli, "hello world")
	if err != nil {
		log.Panic("failed to deploy hello: ", err)
	}
	fmt.Println("addr:", addr.Hex())
	fmt.Println("tx:", tx.Hash())
	msg, err := instance.GetMsg(nil)
	if err != nil {
		log.Panic("failed to getMsg: ", err)
	}
	fmt.Println("msg: ", msg)
}

func main() {
	//1、创建客户端
	cli, err := ethclient.Dial("http://localhost:8541")
	if err != nil {
		log.Panic("failed to dial", err)
	}

	//2.身份准备
	keyin := strings.NewReader(keyinfo)
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		log.Panic("failed to get chainID", err)
	}

	// auth, err := bind.NewTransactor(keyin, "123456")
	auth, err := bind.NewTransactorWithChainID(keyin, "123456", chainID)
	if err != nil {
		log.Panic("auth create failed:", err)
	}
	fmt.Println("chainID: ", chainID)
	auth.GasLimit = 3000000

	//3、部署合约
	//auth *bind.TransactOpts, backend bind.ContractBackend, _msg string
	// addr, tx, instance, err := DeployHello(auth, cli, "hello world")
	instance, err := NewHello(common.HexToAddress("0xEA25C04416a59835486E0bD29eeD251fEE9D630E"), cli)
	if err != nil {
		log.Panic("failed to deploy hello: ", err)
	}

	msg, err := instance.GetMsg(nil)
	if err != nil {
		log.Panic("failed to getMsg: ", err)
	}
	fmt.Println("msg: ", msg)

	tx, err := instance.SetMsg(auth, "fuck world")
	fmt.Println(tx.Hash())

	msg, err = instance.GetMsg(nil)
	if err != nil {
		log.Panic("failed to getMsg: ", err)
	}
	fmt.Println("msg: ", msg)
}
