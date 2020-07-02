package app

import (
	"github.com/jinzhu/gorm"
	deposit "testPipelines/app/v1/deposit/model"
	"testPipelines/app/v1/models"
	redeem "testPipelines/app/v1/redeem/model"
)


func setTestDatabaseData(db *gorm.DB) {
	db.Delete(&deposit.Deposit{})
	db.Delete(&models.TotalVolume{})
	db.Delete(&models.Amount{})
	db.Delete(&models.Profit{})
	db.Delete(&redeem.Redeem{})
	
	db.AutoMigrate(&deposit.Deposit{})
	db.AutoMigrate(&models.TotalVolume{})
	db.AutoMigrate(&models.Amount{})
	db.AutoMigrate(&models.Profit{})
	db.AutoMigrate(&redeem.Redeem{})

	deposit1 := &deposit.Deposit{
		Actor: "hvrlk_1",
		IndexId: "1",
		TotalVolume: []models.TotalVolume{
			{TotalVolume: "deposit_volume1_1", IndexId: "1"},
			{TotalVolume: "deposit_volume2_1", IndexId: "1"},
		},
		Amount: []models.Amount{
			{Amount: "deposit_amount1_1", IndexId: "1"},
			{Amount: "deposit_amount2_1", IndexId: "1"},
		},
		Profit: []models.Profit{
			{Profit: "deposit_profit1_1", IndexId: "1"},
			{Profit: "deposit_profit2_1", IndexId: "1"},
		},
	}
	db.Create(&deposit1)

	redeem1 := &redeem.Redeem{
		Actor: "hvrlk_1",
		IndexId: "1",
		TotalVolume: []models.TotalVolume{
			{TotalVolume: "volume1_1", IndexId: "2"},
			{TotalVolume: "volume2_1", IndexId: "2"},
		},
		Profit: []models.Profit{
			{Profit: "profit1_1", IndexId: "2"},
			{Profit: "profit2_1", IndexId: "2"},
		},
	}
	db.Create(&redeem1)

	deposit2 := &deposit.Deposit{
		Actor: "hvrlk_2",
		IndexId: "2",
		TotalVolume: []models.TotalVolume{
			{TotalVolume: "deposit_volume1_2", IndexId: "3"},
			{TotalVolume: "deposit_volume2_2", IndexId: "3"},
		},
		Amount: []models.Amount{
			{Amount: "deposit_amount1_2", IndexId: "3"},
			{Amount: "deposit_amount2_2", IndexId: "3"},
		},
		Profit: []models.Profit{
			{Profit: "deposit_profit1_2", IndexId: "3"},
			{Profit: "deposit_profit2_2", IndexId: "3"},
		},
	}
	db.Create(&deposit2)

	redeem2 := &redeem.Redeem{
		Actor: "hvrlk_2",
		IndexId: "2",
		TotalVolume: []models.TotalVolume{
			{TotalVolume: "redeem_volume1_2", IndexId: "4"},
			{TotalVolume: "redeem_volume2_2", IndexId: "4"},
		},
		Profit: []models.Profit{
			{Profit: "redeem_profit1_2", IndexId: "4"},
			{Profit: "redeem_profit2_2", IndexId: "4"},
		},
	}
	db.Create(&redeem2)
}
