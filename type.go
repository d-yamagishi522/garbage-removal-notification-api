package main

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

// User firestoreからのデータを整形する用
type User struct {
	ID   string
	Days []Day
}

// Day User struct用
type Day struct {
	Str     string `json:"str"`
	Weekday string `json:"weekday"`
}

// Documents firestoreのdoc取得用
type Documents struct {
	Days []Day `json:"days"`
}

// UserGarbage ユーザーとゴミの紐付き
type UserGarbage struct {
	ID   string
	Text string
}
