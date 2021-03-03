package main

import (
	"bytes"
	"encoding/json"
	mid "garbage-removal-notification-api/middleware"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(mid.CORSMiddleware()))
	e.POST("/sendMessage", sendMessage)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// ライン送信のメインの処理
func sendMessage(ctx echo.Context) error {
	// データを取得
	users := getGarbages()
	// リクエスト用に整形
	bodys := formatGarbages(users)

	// リクエスト処理
	for _, body := range bodys {
		// クライアント作成
		client := &http.Client{}
		// タイムアウトを設定
		client.Timeout = time.Second * 15

		// byteに変更
		jsonValue, _ := json.Marshal(body)
		// リクエストを作成
		req, _ := http.NewRequest("POST", "https://api.line.me/v2/bot/message/multicast", bytes.NewBuffer(jsonValue))

		// ヘッダーを追加
		header := http.Header{}
		header.Set("Content-Length", "10000")
		header.Add("Content-Type", "application/json")
		header.Add("Authorization", "Bearer "+os.Getenv("LINE_ACCESS_TOKEN"))
		req.Header = header

		// リクエストを実行
		resp, err := client.Do(req)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		defer resp.Body.Close()
	}

	return ctx.NoContent(http.StatusNoContent)
}
