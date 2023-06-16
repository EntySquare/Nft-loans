package app

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math/big"
	"nft-loans/config"
	"nft-loans/contracts"
	"nft-loans/database"
	"nft-loans/model"
	"nft-loans/pkg"
	"nft-loans/routing/types"
	"strings"
	"time"
)

func Deposit(c *fiber.Ctx) error {
	//reqParams := types.DepositNgtReq{}
	//err := c.BodyParser(&reqParams)
	//if err != nil {
	//	return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	//}
	//userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	//err = database.DB.Transaction(func(tx *gorm.DB) error {
	//	acc := model.Account{}
	//	acc.UserId = userId
	//	err = acc.GetByUserId(database.DB)
	//	if err != nil {
	//		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), ""))
	//	}
	//	acc.FrozenBalance += reqParams.Num
	//	err = acc.UpdateAccount(tx)
	//	if err != nil {
	//		return err
	//	}
	//	tt := time.Now()
	//	acf := model.AccountFlow{
	//		AccountId:       acc.ID,
	//		Num:             reqParams.Num,
	//		Chain:           reqParams.Chain,
	//		Address:         reqParams.Address,
	//		Hash:            reqParams.Hash,
	//		AskForTime:      &tt,
	//		AchieveTime:     nil,
	//		TransactionType: "1",
	//		Flag:            "1",
	//	}
	//	err = acf.InsertNewAccountFlow(tx)
	//	if err != nil {
	//		return err
	//	}
	//	txs := model.Transactions{
	//		Hash:      reqParams.Hash,
	//		Status:    "0",
	//		ChainName: reqParams.Chain,
	//		Flag:      "1",
	//	}
	//	err := txs.InsertNewTransactions(tx)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), ""))
	//}

	return c.JSON(pkg.SuccessResponse(""))
}
func Withdraw(c *fiber.Ctx) error {
	reqParams := types.WithdrawNgtReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	hash := contracts.TransferFrom(common.HexToAddress("Manager"), common.HexToAddress(reqParams.Address), big.NewInt(reqParams.Num), reqParams.Chain)
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		acc := model.Account{}
		acc.UserId = userId
		err = acc.GetByUserId(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), ""))
		}
		acc.Balance -= reqParams.Num
		err = acc.UpdateAccount(tx)
		if err != nil {
			return err
		}
		tt := time.Now()
		acf := model.AccountFlow{
			AccountId:       acc.ID,
			Num:             float64(reqParams.Num),
			Chain:           reqParams.Chain,
			Address:         reqParams.Address,
			Hash:            hash,
			AskForTime:      &tt,
			AchieveTime:     nil,
			TransactionType: "2",
			Flag:            "1",
		}
		err = acf.InsertNewAccountFlow(tx)
		if err != nil {
			return err
		}
		txs := model.Transactions{
			Hash:      reqParams.Hash,
			Status:    "0",
			ChainName: reqParams.Chain,
			Flag:      "1",
		}
		err := txs.InsertNewTransactions(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}

func CheckHashApi(c *fiber.Ctx) error {
	var (
		db        = database.DB
		reqParams = types.CheckHashApiReq{}
	)
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}

	err = db.Create(&model.Transactions{
		Hash:      reqParams.Hash,
		Status:    "1",
		ChainName: "poly",
		Flag:      "1",
	}).Error
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Create Transactions err", ""))
	}

	go func() {
		from, to, _, f, err := contracts.CheckHash(reqParams.Hash)
		if err != nil {
			return
		}
		fromStr := strings.ToLower(from)
		toStr := strings.ToLower(to)
		adminStr := strings.ToLower(config.Config("CONTRACT_ADDRESS"))
		//是我方收钱
		if toStr != adminStr {
			return
		}
		user := model.User{WalletAddress: fromStr}
		if err = user.GetByWalletAddress(db); err != nil {
			return
		}
		//加余额
		err = db.Model(&model.Account{}).
			Where("id = ?", user.ID).
			Update("balance", gorm.Expr("balance + ?", f)).Error
		if err != nil {
			panic(err)
		}
		db.Model(&model.Transactions{}).Where("hash = ?", reqParams.Hash).
			Updates(map[string]interface{}{"status": "2", "from_address": fromStr})
		//gorm.Model
		acc := model.Account{}
		acc.UserId = user.ID
		err = acc.GetByUserId(database.DB)
		if err != nil {
			panic(err)
		}
		now := time.Now()
		acf := model.AccountFlow{
			AccountId:       acc.ID,
			Num:             f,
			Chain:           "poly",
			Address:         fromStr,
			Hash:            reqParams.Hash,
			AskForTime:      &now,
			AchieveTime:     nil,
			TransactionType: "1",
			Flag:            "1",
		}
		database.DB.Create(&acf)
	}()

	return c.JSON(pkg.SuccessResponse(""))
}
