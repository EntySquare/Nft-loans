package app

import (
	"github.com/gofiber/fiber/v2"
	"nft-loans/config"
	"nft-loans/pkg"
	"nft-loans/routing/types"
)

func Deposit(c *fiber.Ctx) error {
	reqParams := types.LoginAndRegisterReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}

	returnT := pkg.RandomString(64)

	return c.JSON(pkg.SuccessResponse(returnT))
}
func Withdraw(c *fiber.Ctx) error {
	reqParams := types.LoginAndRegisterReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	returnT := pkg.RandomString(64)

	return c.JSON(pkg.SuccessResponse(returnT))
}
