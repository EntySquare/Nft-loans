package routing

import (
	"github.com/gofiber/fiber/v2"
	"nft-loans/routing/app"
	intcpt "nft-loans/routing/intercept"
)

func Setup(f *fiber.App) {
	appApi := f.Group("/app")

	AppUserSetUp(appApi)
	AppNftSetUp(appApi)

}

func AppUserSetUp(appApi fiber.Router) {
	appApi.Post("/login", app.LoginAndRegister) //login
}
func AppNftSetUp(appApi fiber.Router) {
	appApi.Post("/nft/deposit", intcpt.AuthApp(), app.Deposit)
	appApi.Post("/nft/withdraw", intcpt.AuthApp(), app.Withdraw)
	appApi.Post("/nft/approve", intcpt.AuthApp(), app.Approve)
}
