package app

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"nft-loans/config"
	"nft-loans/database"
	"nft-loans/model"
	"nft-loans/pkg"
	"nft-loans/routing/types"
	"strconv"
	"time"
)

func PledgeNft(c *fiber.Ctx) error {
	reqParams := types.PledgeNgtReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	tt := time.Now()
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		atoi, err := strconv.Atoi(reqParams.Duration)
		if err != nil {
			return err
		}
		tf := tt.AddDate(0, 0, atoi)
		co := model.Covenant{
			NFTName:            "TEST",
			PledgeId:           reqParams.NftId,
			ChainName:          reqParams.Chain,
			Duration:           reqParams.Duration,
			Hash:               reqParams.Hash,
			InterestRate:       0.6,
			AccumulatedBenefit: 0,
			PledgeFee:          0,
			ReleaseFee:         0,
			StartTime:          &tt,
			ExpireTime:         &tf,
			NFTReleaseTime:     &tt,
			Flag:               "1",
			OwnerId:            userId,
		}
		err = co.InsertNewCovenant(tx)
		if err != nil {
			return err
		}
		txs := model.Transactions{
			Hash:      reqParams.Hash,
			Status:    "1",
			ChainName: reqParams.Chain,
			Flag:      "1",
		}
		err = txs.InsertNewTransactions(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "pledge error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
