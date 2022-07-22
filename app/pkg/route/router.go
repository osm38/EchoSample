package route

import (
	"github.com/labstack/echo/v4"
	"echosample/pkg/handle"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/gorilla/sessions"
	// "github.com/labstack/echo-contrib/session"
	"echosample/pkg/authenticate"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// ロガー設定（デフォルトは標準出力）
	e.Use(middleware.Logger())
	// panic（システム例外的なもの）が起こったときにサーバを落とさずにエラーレスポンスを返す。
	e.Use(middleware.Recover())
	// セッション管理
	// e.Use(session.Middleware(sessions.NewCookieStore([]byte("GoSampleSession"))))
	// カスタム認証
	e.Use(authenticate.SimpleAuthWithConfig(authenticate.SimpleAuthConfig{}))

	// 静的コンテンツフォルダの設定
	e.Static("/assets", "web/assets")

	// ルーティング設定
	// e.File("/", "web/template/index.html")
	e.POST("/", handle.PostRoot)
	e.GET("/hello", handle.GetHelloWorld)
	e.GET("/users/:id", handle.GetUserID)
	e.POST("/user", handle.GetUser)
	e.POST("/login", handle.DoLogin)
	e.GET("/sample", handle.AjaxSample)

	return e
}