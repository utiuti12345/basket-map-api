package main

import (
	"basket-map-api/infrastructure/postgres"
	"basket-map-api/infrastructure/web"
	"basket-map-api/interface/controller"
	"basket-map-api/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	db := postgres.SetupDB()
	if db == nil {
		log.Error("DB接続に失敗しました")
		return
	}

	r := postgres.NewParkRepository(db)
	u := usecase.NewParkUsecase(r)
	c := controller.NewParkController(u)

	// ルーティングの設定
	e := echo.New()
	if e == nil {
		log.Error("echoの初期化に失敗しました")
		return
	}

	web.NewRouter(e, c)

	if err := e.Start(":8080"); err != nil {
		log.Error("サーバーの起動に失敗しました", "err", err.Error())
		return
	}
}
