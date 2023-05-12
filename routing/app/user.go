package app

import (
	"github.com/gofiber/fiber/v2"
	"nft-loans/config"
	"nft-loans/database"
	"nft-loans/model"
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
