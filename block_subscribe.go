package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main(){
	client,err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil{
		log.Fatal(err)
	}
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for{
		select{
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

		block, err := client.BlockByHash(context.Background(), header.Hash())
		if err != nil{
			log.Fatal(err)
		}

		fmt.Println(block.Hash().Hex())
		fmt.Println(block.Number().Uint64()) // 3477413
		fmt.Println(block.Time()) // 1529525947
		fmt.Println(block.Nonce()) // 130524141876765836
		fmt.Println(len(block.Transactions()))



		}
	}



}