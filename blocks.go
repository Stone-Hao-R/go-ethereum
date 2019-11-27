package main
import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)


func main(){
	client , err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil{
		log.Fatal(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(header.Number.String()) // 5671744

	blockNumber := big.NewInt(9002950)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number().Uint64()) // 5671744
	fmt.Println(block.Time()) // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex()) // 0x9e8751ebb5069389b855bba72d94902cc38504266149 8a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions())) // 144

	count, err := client.TransactionCount(context.Background(),block.Hash())
	if err != nil {
		log.Fatal(err)}
	fmt.Println(count) // 144

	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083c a52ff24b3b65bc9c2
		fmt.Println(tx.Value().String()) // 10000000000000000
		fmt.Println(tx.Gas()) // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce()) // 110644
		fmt.Println(tx.Data()) // []
		fmt.Println(tx.To().Hex()) // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		chainID, err := client.NetworkID(context.Background())
		fmt.Println("--------")

		if err != nil{
			log.Fatal(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err != nil{

			fmt.Println(msg)
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err!=nil{
			log.Fatal(err)

		}

		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		break; //打印一条看看
	}


}