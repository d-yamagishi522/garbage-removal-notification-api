package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

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

// firestoreから取得したデータをリクエスト用に整形する処理
func formatGarbages(users []User) []Body {
	// 曜日を取得
	t := time.Now()
	wd := t.Weekday().String()

	// ユーザーとゴミのメッセージを1対1に整形
	garbages := []UserGarbage{}
	for _, user := range users {
		for _, item := range user.Days {
			if item.Weekday == wd && item.Garbage != "" {
				obj := UserGarbage{
					ID:   user.ID,
					Text: item.Garbage + "回収の日",
				}
				garbages = append(garbages, obj)
			}
		}
	}

	// リクエストの形に整形
	bodys := []Body{}
	for _, garbage := range garbages {
		if len(bodys) == 0 {
			to := []string{garbage.ID}
			messages := []Message{}
			message := Message{
				Type: "text",
				Text: garbage.Text,
			}
			messages = append(messages, message)
			body := Body{
				To:       to,
				Messages: messages,
			}
			bodys = append(bodys, body)
		} else {
			existed := false
			for _, b := range bodys {
				if b.Messages[0].Text == garbage.Text {
					existed = true
					b.To = append(b.To, garbage.ID)
				}
			}
			if !existed {
				to := []string{garbage.ID}
				messages := []Message{}
				message := Message{
					Type: "text",
					Text: garbage.Text,
				}
				messages = append(messages, message)
				body := Body{
					To:       to,
					Messages: messages,
				}
				bodys = append(bodys, body)
			}
		}
	}
	return bodys
}
