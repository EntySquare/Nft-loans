package app

import (
	"github.com/gofiber/fiber/v2"
	"nft-loans/config"
	"nft-loans/database"
	"nft-loans/model"
	"nft-loans/model/api"
	"nft-loans/pkg"
	"nft-loans/routing/types"
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
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "user get by phone error", ""))
	}
	returnT := pkg.RandomString(64)

	err = user.InsertNewUser(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "注册失败"))
	}
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
	reqParams := types.LoginAndRegisterReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	user := model.User{
		Flag: "1",
	}
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "user get by phone error", ""))
	}
	returnT := pkg.RandomString(64)

	err = user.InsertNewUser(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "注册失败"))
	}
	return c.JSON(pkg.SuccessResponse(returnT))
}
func MyBenefit(c *fiber.Ctx) error {
	reqParams := types.LoginAndRegisterReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	user := model.User{
		Flag: "1",
	}
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "user get by phone error", ""))
	}
	returnT := pkg.RandomString(64)

	err = user.InsertNewUser(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, err.Error(), "注册失败"))
	}
	return c.JSON(pkg.SuccessResponse(returnT))
}
