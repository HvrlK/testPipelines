package app

import (
	"github.com/jinzhu/gorm"
	deposit "testPipelines/app/v1/deposit/model"
	"testPipelines/app/v1/models"
	redeem "testPipelines/app/v1/redeem/model"
)


func setTestDatabaseData(db *gorm.DB) {
	db.AutoMigrate(&deposit.Deposit{})
	db.AutoMigrate(&models.TotalVolume{})
	db.AutoMigrate(&models.Amount{})
	db.AutoMigrate(&models.Profit{})
	db.AutoMigrate(&redeem.Redeem{})

	db.Delete(&deposit.Deposit{})
	db.Delete(&models.TotalVolume{})
	db.Delete(&models.Amount{})
	db.Delete(&models.Profit{})
	db.Delete(&redeem.Redeem{})

	deposit1Uid := deposit.MadeUid("hvrlk_1", "1")
	deposit1 := &deposit.Deposit{
		Uid: deposit1Uid,
		Actor: "hvrlk_1",
		IndexId: "1",
		TotalVolume: []models.TotalVolume{
			{TotalVolume: "deposit_volume1_1", Uid: deposit1Uid},
			{TotalVolume: "deposit_volume2_1", Uid: deposit1Uid},
		},
		Amount: []models.Amount{
			{Amount: "deposit_amount1_1", Uid: deposit1Uid},
			{Amount: "deposit_amount2_1", Uid: deposit1Uid},
		},
		Profit: []models.Profit{
			{Profit: "deposit_profit1_1", Uid: deposit1Uid},
			{Profit: "deposit_profit2_1", Uid: deposit1Uid},
		},
	}
	db.Create(&deposit1)

	redeem1Uid := redeem.MadeUid("hvrlk_1", "2")
	redeem1 := &redeem.Redeem{
		Uid: redeem1Uid,
		Actor: "hvrlk_1",
		IndexId: "2",
		TotalVolume: []models.TotalVolume{
			{TotalVolume: "volume1_1", Uid: redeem1Uid},
			{TotalVolume: "volume2_1", Uid: redeem1Uid},
		},
		Profit: []models.Profit{
			{Profit: "profit1_1", Uid: redeem1Uid},
			{Profit: "profit2_1", Uid: redeem1Uid},
		},
	}
	db.Create(&redeem1)

	deposit2Uid := deposit.MadeUid("hvrlk_2", "3")
	deposit2 := &deposit.Deposit{
		Uid: deposit2Uid,
		Actor: "hvrlk_2",
		IndexId: "3",
		TotalVolume: []models.TotalVolume{
			{TotalVolume: "deposit_volume1_2", Uid: deposit2Uid},
			{TotalVolume: "deposit_volume2_2", Uid: deposit2Uid},
		},
		Amount: []models.Amount{
			{Amount: "deposit_amount1_2", Uid: deposit2Uid},
			{Amount: "deposit_amount2_2", Uid: deposit2Uid},
		},
		Profit: []models.Profit{
			{Profit: "deposit_profit1_2", Uid: deposit2Uid},
			{Profit: "deposit_profit2_2", Uid: deposit2Uid},
		},
	}
	db.Create(&deposit2)

	redeem2Uid := redeem.MadeUid("hvrlk_2", "4")
	redeem2 := &redeem.Redeem{
		Uid: redeem2Uid,
		Actor: "hvrlk_2",
		IndexId: "4",
		TotalVolume: []models.TotalVolume{
			{TotalVolume: "redeem_volume1_2", Uid: redeem2Uid},
			{TotalVolume: "redeem_volume2_2", Uid: redeem2Uid},
		},
		Profit: []models.Profit{
			{Profit: "redeem_profit1_2", Uid: redeem2Uid},
			{Profit: "redeem_profit2_2", Uid: redeem2Uid},
		},
	}
	db.Create(&redeem2)
}
