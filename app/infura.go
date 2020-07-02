package app

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

func ConnectToInfura() {
	infuraApiKey := os.Getenv("INFURA_API_KEY")

	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/" + infuraApiKey)
	if err != nil {
		log.Fatal(err)
	}

	go handleNewBlock(client)
	go subscribeLogs(client)
}

func subscribeLogs(client *ethclient.Client)  {
	contractAddress := common.HexToAddress("0xad6d458402f60fd3bd25163575031acdce07538d")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}

func handleNewBlock(client *ethclient.Client) {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Error: ", err)
			continue
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())

			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

			fmt.Println("Hex: ", block.Hash().Hex())
			fmt.Println("Number: ", block.Number().Uint64())
			fmt.Println("Time: ", block.Time())
		}
	}
}
