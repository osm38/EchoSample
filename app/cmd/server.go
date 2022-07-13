package main

import (
	"echosample/pkg/route"
	"echosample/pkg/db/repository"
)

func main(){
	repository.FindUser()

	// ルーティング設定
	router := route.NewRouter()
	// ポート1323で待ち受け
	router.Logger.Fatal(router.Start(":1323"))
}
