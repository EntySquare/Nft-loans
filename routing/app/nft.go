package app

import (
	"github.com/gofiber/fiber/v2"
	"nft-loans/config"
	"nft-loans/pkg"
	"nft-loans/routing/types"
)

func PledgeNft(c *fiber.Ctx) error {
	reqParams := types.PledgeNgtReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "parser error", ""))
	}
	println(reqParams.NftId)
	println(reqParams.Chain)
	println(reqParams.Duration)
	println(reqParams.Hash)
	return c.JSON(pkg.SuccessResponse(""))
}
