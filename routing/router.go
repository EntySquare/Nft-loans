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
	AppNftSetUp(appApi)

}

func AppUserSetUp(appApi fiber.Router) {
	appApi.Post("/login", app.LoginAndRegister)                      //login
	appApi.Post("/myBenefit", intcpt.AuthApp(), app.MyBenefit)       //login
	appApi.Post("/myNgt", intcpt.AuthApp(), app.MyNgt)               //login
	appApi.Post("/myInvestment", intcpt.AuthApp(), app.MyInvestment) //login
}
func AppWalletSetUp(appApi fiber.Router) {
	appApi.Post("/wallet/deposit", intcpt.AuthApp(), app.Deposit)
	appApi.Post("/wallet/withdraw", intcpt.AuthApp(), app.Withdraw)
}
func AppNftSetUp(appApi fiber.Router) {
	appApi.Post("/nft/approve", intcpt.AuthApp(), app.Approve)
}
