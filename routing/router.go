package routing

import (
	"github.com/gofiber/fiber/v2"
	"nft-loans/routing/app"
	intcpt "nft-loans/routing/intercept"
)

func Setup(f *fiber.App) {
	appApi := f.Group("/app")

	AppUserSetUp(appApi)
	AppWalletSetUp(appApi)

}

func AppUserSetUp(appApi fiber.Router) {
	appApi.Post("/login", app.LoginAndRegister) //login
}
func AppWalletSetUp(appApi fiber.Router) {
	appApi.Post("/wallet/deposit", intcpt.AuthApp(), app.Deposit)
	appApi.Post("/wallet/withdraw", intcpt.AuthApp(), app.Withdraw)
}
func AppNftSetUp(appApi fiber.Router) {
	appApi.Post("/nft/approve", intcpt.AuthApp(), app.Approve)
}
