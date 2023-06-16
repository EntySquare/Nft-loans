package app

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"nft-loans/config"
	"nft-loans/database"
	"nft-loans/model"
	"nft-loans/model/api"
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

	//如果存在了则不能再次插入

	cc := model.Covenant{}
	database.DB.Model(&model.Covenant{}).
		Where("owner_id = ? and flag = '1' and pledge_id = ?",
			userId, reqParams.NftId).Take(&cc)
	if cc.ID != 0 { //有数据
		return c.JSON(pkg.SuccessResponse(""))
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		nftId, err := strconv.Atoi(reqParams.NftId)
		if err != nil {
			return err
		}
		nftName, _, interestRate := api.GetInterestRate(nftId)

		atoi, err := strconv.Atoi(reqParams.Duration)
		if err != nil {
			return err
		}
		tf := tt.AddDate(0, 0, atoi)
		co := model.Covenant{
			NFTName:            nftName,
			PledgeId:           reqParams.NftId,
			ChainName:          reqParams.Chain,
			Duration:           reqParams.Duration,
			Hash:               reqParams.Hash,
			InterestRate:       interestRate,
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
		go func() {
			i := 0
			for {
				if i >= 5 {
					return
				}
				api.SelectChainData(database.DB, userId)
				time.Sleep(time.Second * 30)
				i++
			}
		}()
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "pledge error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
func CancelCovenant(c *fiber.Ctx) error {
	reqParams := types.CancelCovenantReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	tt := time.Now()
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)

	//如果存在了则不能再次插入

	cc := model.Covenant{}
	database.DB.Model(&model.Covenant{}).
		Where("owner_id = ? and flag = '1' and pledge_id = ?",
			userId, reqParams.NftId).Take(&cc)
	if cc.ID == 0 { //没有数据
		return c.JSON(pkg.SuccessResponse(""))
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		tf := tt.AddDate(0, 0, 7)
		cc.ExpireTime = &tf
		cc.Flag = "0"
		err = cc.UpdateCovenant(tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "cancel error", ""))
	}
	return c.JSON(pkg.SuccessResponse(""))
}
