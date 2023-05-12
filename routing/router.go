package routing

import (
	"github.com/gofiber/fiber/v2"
	"nft-loans/routing/app"
)

func Setup(f *fiber.App) {
	appApi := f.Group("/app")

	AppProductSetUp(appApi)

}

// AppSetUp
func AppProductSetUp(appApi fiber.Router) {
	appApi.Post("/login", app.LoginAndRegister) //获取btc涨幅和usdt价
}
