package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type TotalVolume struct {
	gorm.Model
	TotalVolume string
	IndexId string
}

type Amount struct {
	gorm.Model
	Amount string
	IndexId string
}

type Profit struct {
	gorm.Model
	Profit string
	IndexId string
}

type Deposit struct {
	gorm.Model
	Actor string
	IndexId string
	TotalVolume []TotalVolume `gorm:"foreignkey:IndexId;association_foreignkey:IndexId"`
	Amount []Amount           `gorm:"foreignkey:IndexId;association_foreignkey:IndexId"`
	Profit []Profit           `gorm:"foreignkey:IndexId;association_foreignkey:IndexId"`
}

type Redeem struct {
	gorm.Model
	Actor string
	IndexId string
	TotalVolume []TotalVolume `gorm:"foreignkey:IndexId;association_foreignkey:IndexId"`
	Profit []Profit           `gorm:"foreignkey:IndexId;association_foreignkey:IndexId"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	testDatabase(db)

	r := gin.Default()

	r.GET("/profit/:sender/:indexId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"profit": "100",
		})
	})

	r.GET("/deposit/:actor", func (c *gin.Context) {
		result := &Deposit{}
		db.Preload("TotalVolume", &result.TotalVolume).
			Preload("Amount", &result.Amount).
			Preload("Profit", &result.Profit).
			Where("Actor = ?", c.Param("actor")).
			First(result)
		c.JSON(200, result)
	})

	r.GET("/redeem/:actor", func (c *gin.Context) {
		result := &Deposit{}
		db.Preload("TotalVolume", &result.TotalVolume).
			Preload("Profit", &result.Profit).
			Where("Actor = ?", c.Param("actor")).
			First(result)
		c.JSON(200, result)
	})

	infuraApiKey := os.Getenv("INFURA_API_KEY")

	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/" + infuraApiKey)
	if err != nil {
		log.Fatal(err)
	}

	go handleNewBlock(client)
	go subscribeLogs(client)

	_ = r.Run(":" + port)

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

func testDatabase(db *gorm.DB) {
	db.AutoMigrate(&Deposit{})
	db.AutoMigrate(&TotalVolume{})
	db.AutoMigrate(&Amount{})
	db.AutoMigrate(&Profit{})
	db.AutoMigrate(&Redeem{})

	deposit1 := Deposit{
		Actor: "hvrlk_1",
		IndexId: "1",
		TotalVolume: []TotalVolume{
			{TotalVolume: "deposit_volume1_1", IndexId: "1"},
			{TotalVolume: "deposit_volume2_1", IndexId: "1"},
		},
		Amount: []Amount{
			{Amount: "deposit_amount1_1", IndexId: "1"},
			{Amount: "deposit_amount2_1", IndexId: "1"},
		},
		Profit: []Profit{
			{Profit: "deposit_profit1_1", IndexId: "1"},
			{Profit: "deposit_profit2_1", IndexId: "1"},
		},
	}
	db.Create(&deposit1)

	redeem1 := Redeem{
		Actor: "hvrlk_1",
		IndexId: "1",
		TotalVolume: []TotalVolume{
			{TotalVolume: "volume1_1", IndexId: "2"},
			{TotalVolume: "volume2_1", IndexId: "2"},
		},
		Profit: []Profit{
			{Profit: "profit1_1", IndexId: "2"},
			{Profit: "profit2_1", IndexId: "2"},
		},
	}
	db.Create(&redeem1)

	deposit2 := &Deposit{
		Actor: "hvrlk_2",
		IndexId: "2",
		TotalVolume: []TotalVolume{
			{TotalVolume: "deposit_volume1_2", IndexId: "3"},
			{TotalVolume: "deposit_volume2_2", IndexId: "3"},
		},
		Amount: []Amount{
			{Amount: "deposit_amount1_2", IndexId: "3"},
			{Amount: "deposit_amount2_2", IndexId: "3"},
		},
		Profit: []Profit{
			{Profit: "deposit_profit1_2", IndexId: "3"},
			{Profit: "deposit_profit2_2", IndexId: "3"},
		},
	}
	db.Create(&deposit2)

	redeem2 := &Redeem{
		Actor: "hvrlk_2",
		IndexId: "2",
		TotalVolume: []TotalVolume{
			{TotalVolume: "redeem_volume1_2", IndexId: "4"},
			{TotalVolume: "redeem_volume2_2", IndexId: "4"},
		},
		Profit: []Profit{
			{Profit: "redeem_profit1_2", IndexId: "4"},
			{Profit: "redeem_profit2_2", IndexId: "4"},
		},
	}
	db.Create(&redeem2)
}