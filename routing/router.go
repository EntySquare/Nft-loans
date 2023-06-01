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
	appApi.Post("/login", app.LoginAndRegister)                                //login
	appApi.Post("/myCovenantFlow", intcpt.AuthApp(), app.MyCovenantFlow)       //login
	appApi.Post("/myNgt", intcpt.AuthApp(), app.MyNgt)                         //login
	appApi.Post("/myInvestment", intcpt.AuthApp(), app.MyInvestment)           //login
	appApi.Post("/getInviteeInfo", intcpt.AuthApp(), app.GetInviteeInfo)       //login
	appApi.Post("/getCovenantDetail", intcpt.AuthApp(), app.GetCovenantDetail) //login
}
func AppWalletSetUp(appApi fiber.Router) {
	appApi.Post("/wallet/deposit", intcpt.AuthApp(), app.Deposit)
	appApi.Post("/wallet/withdraw", intcpt.AuthApp(), app.Withdraw)
}
func AppNftSetUp(appApi fiber.Router) {
	appApi.Post("/nft/pledgeNft", intcpt.AuthApp(), app.PledgeNft)
}
