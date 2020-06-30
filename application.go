package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	r := gin.Default()

	r.GET("/profit/:sender/:indexId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"profit": "100",
		})
	})

	infuraApiKey := os.Getenv("INFURA_API_KEY")

	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/" + infuraApiKey)
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	go handleNewBlock(client, headers, sub)

	_ = r.Run(":" + port)

}

func handleNewBlock(client *ethclient.Client, headers chan *types.Header, sub ethereum.Subscription) {
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