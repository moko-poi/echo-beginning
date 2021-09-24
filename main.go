package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	//"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echoインスタンス作成(*http.Serverやnet.Listener)
	e := echo.New()

	// HTTPメソッドとリクエストを受け取るURLパターン､それに対応するためのハンドラを登録(http.HandleFunc)
	e.GET("/", func(c echo.Context) error  {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Logger.Fatal(e.Start(":1323"))
}


// 補足　echo.Context　現在のHTTPリクエスト状況　（http.ResponseWriterや*http.Request　ざまざまな情報を保存）
// context.string ステータスコードとともに文字列をレスポンスに書き込むためのメソッド
// (*Echo) Start HTTPサーバーを開始するためのメソッド　（指定されたアドレスへのリクエストを待ち受け開始します　http.ListenAndServeなど）