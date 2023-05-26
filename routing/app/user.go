package app

import (
	"github.com/gofiber/fiber/v2"
	"nft-loans/config"
	"nft-loans/database"
	"nft-loans/model"
	"nft-loans/model/api"
	"nft-loans/pkg"
	"nft-loans/routing/types"
	"strconv"
	"strings"
	"time"
)

// LoginAndRegister 登录注册
func LoginAndRegister(c *fiber.Ctx) error {
	reqParams := types.LoginAndRegisterReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	user := model.User{
		Flag: "1",
	}
	user.WalletAddress = reqParams.WalletAddress
	returnT := ""
	err = user.GetByWalletAddress(database.DB)
	if err != nil {
		if !strings.Contains(err.Error(), "record not found") {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "user get by addresss error", ""))
		}
		user.Level = 0
		user.PledgeCount = 0
		returnT = pkg.RandomString(64)
		user.Token = returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
		err = user.InsertNewUser(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "注册失败"))
		}
		user.RecommendId = reqParams.RecommendId
		err := user.GetByWalletAddress(database.DB)
		if err != nil {
			return err
		}
		var acc = model.Account{
			UserId:        user.ID,
			Balance:       0,
			FrozenBalance: 0,
			Flag:          "1",
		}
		err = acc.InsertNewAccount(database.DB)
		if err != nil {
			return err
		}
	}
	returnT = strings.Split(user.Token, ":")[0]
	c.Locals(config.LOCAL_TOKEN, returnT)
	return c.JSON(pkg.SuccessResponse(returnT))
}
func MyInvestment(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	user := model.User{}
	user.ID = userId
	err := user.GetById(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询用户失败"))
	}
	data := types.MyInvestmentResp{
		InvestmentUsers: make([]types.InvestmentUserInfo, 0),
	}
	data.UID = user.UID
	data.InvestmentAddress = user.InvestmentAddress
	data.Level = user.Level
	data.InvestmentCount = int64(len(api.UserTree[userId].Branch))
	data.AccumulatedPledgeCount = api.GetBranchAccumulatedPledgeCount(userId)
	for _, branch := range api.UserTree[userId].Branch {
		in := types.InvestmentUserInfo{}
		in.UID = api.UserTree[branch].UID
		in.Level = api.UserTree[branch].Level
		in.PledgeCount = api.UserTree[branch].PledgeCount
		data.InvestmentUsers = append(data.InvestmentUsers, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
func MyNgt(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	acc := model.Account{}
	acc.UserId = userId
	err := acc.GetByUserId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户失败"))
	}
	benefit, err := getBenefit(acc)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询收益失败"))
	}
	data := types.MyNgtResp{
		BenefitInfo:  benefit,
		Transactions: make([]types.TransactionInfo, 0),
	}
	af := model.AccountFlow{}
	af.AccountId = acc.ID
	afs, err := af.GetByAccountId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户交易失败"))
	}
	for _, accflow := range afs {
		txs := model.Transactions{}
		txs.Hash = accflow.Hash
		err := txs.GetByHash(database.DB)
		if err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户交易失败"))
		}
		in := types.TransactionInfo{
			Num:             accflow.Num,
			Chain:           accflow.Chain,
			Address:         accflow.Address,
			Hash:            accflow.Hash,
			AskForTime:      accflow.AskForTime,
			AchieveTime:     accflow.AchieveTime,
			TransactionType: accflow.TransactionType,
			Status:          txs.Status,
		}
		data.Transactions = append(data.Transactions, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
func MyCovenantFlow(c *fiber.Ctx) error {
	//userId := c.Locals("user_id")
	userId := c.Locals(config.LOCAL_USERID_UINT).(uint)
	acc := model.Account{}
	acc.UserId = userId
	err := acc.GetByUserId(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户失败"))
	}
	benefit, err := getBenefit(acc)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询收益失败"))
	}
	data := types.MyCovenantFlowResp{
		BenefitInfo: benefit,
		Covenants:   make([]types.CovenantInfo, 0),
	}
	co := model.Covenant{}
	co.OwnerId = acc.UserId
	cos, err := co.SelectMyCovenant(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "查询账户交易失败"))
	}
	for _, coi := range cos {
		in := types.CovenantInfo{
			NFTName:            coi.NFTName,
			PledgeId:           coi.PledgeId,
			ChainName:          coi.ChainName,
			Duration:           coi.Duration,
			Hash:               coi.Hash,
			InterestRate:       coi.InterestRate,
			AccumulatedBenefit: coi.AccumulatedBenefit,
			PledgeFee:          coi.PledgeFee,
			ReleaseFee:         coi.ReleaseFee,
			StartTime:          coi.StartTime,
			ExpireTime:         coi.ExpireTime,
			NFTReleaseTime:     coi.NFTReleaseTime,
			Flag:               coi.Flag,
		}
		data.Covenants = append(data.Covenants, in)

	}
	return c.JSON(pkg.SuccessResponse(data))
}
func getLastDay() int64 {
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -1)
	y, m, d := oldTime.Date()
	date := int64(y*10000 + int(m)*100 + d)
	return date
}
func getBenefit(acc model.Account) (types.BenefitInfo, error) {
	data := types.BenefitInfo{
		LastDayBenefit: 0.0,
	}
	data.Balance = acc.Balance
	co := model.Covenant{}
	co.OwnerId = acc.UserId
	benefits, err := co.GetUserAccumulatedBenefit(database.DB)
	if err != nil {
		return data, err
	}
	data.AccumulatedBenefit = benefits
	cf := model.CovenantFlow{}
	cf.AccountId = acc.ID
	cf.ReleaseDate = getLastDay()
	cfs, err := cf.GetByAccountIdAndReleaseDate(database.DB)
	if err != nil {
		return data, err
	}
	for _, flow := range cfs {
		data.LastDayBenefit += flow.Num
	}
	return data, nil
}
