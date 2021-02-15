package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

// Body struce
type Body struct {
	To       []string  `json:"to"`
	Messages []Message `json:"messages"`
}

// Message struct
type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func sendMessage(ctx echo.Context) error {
	// クライアント作成
	client := &http.Client{}
	// タイムアウトを設定
	client.Timeout = time.Second * 15

	// リクエストボディを作成
	// TODO: ラインユーザーのuidを取得する処理を作成する
	uid := ""
	uids := []string{}
	uids = append(uids, uid)

	messages := []Message{}
	// TODO: firebaseからユーザーごとにメッセージを取得する
	message := Message{
		Type: "text",
		Text: "sample",
	}
	messages = append(messages, message)
	body := Body{
		To:       uids,
		Messages: messages,
	}
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

	return ctx.NoContent(http.StatusNoContent)
}
