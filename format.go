package main

import (
	"time"
)

// firestoreから取得したデータをリクエスト用に整形する処理
func formatGarbages(users []User) []Body {
	// 曜日を取得
	time.Local = time.FixedZone("UTC", 0)
	t := time.Now()
	jst := t.Add(9 * time.Hour)
	wd := jst.Weekday().String()

	// ユーザーとゴミのメッセージを1対1に整形
	garbages := []UserGarbage{}
	for _, user := range users {
		for _, item := range user.Days {
			if item.Weekday == wd && item.Str != "" {
				text := "本日は、" + item.Str + `回収の日です。

通知を停止する場合は以下のリンクから設定してください。
https://grn-line.herokuapp.com/
`
				obj := UserGarbage{
					ID:   user.ID,
					Text: text,
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
