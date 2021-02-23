package main

import (
	"context"
	"encoding/json"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// firestoreからデータを取得する処理
func getGarbages() []User {
	// firebaseの設定
	ctx := context.Background()
	opt := option.WithCredentialsFile("secretkey.json")
	app, _ := firebase.NewApp(context.Background(), nil, opt)
	// firestoreの設定
	client, _ := app.Firestore(ctx)
	// usersコレクションを全取得
	docRefs, _ := client.Collection("users").Where("isNotificated", "==", true).Documents(ctx).GetAll()

	// firestoreから取得したデータを整形
	users := []User{}
	for _, doc := range docRefs {
		obj := doc.Data()
		l := Documents{}
		tmp, _ := json.Marshal(obj)
		_ = json.Unmarshal(tmp, &l)
		u := User{
			ID:   doc.Ref.ID,
			Days: l.Days,
		}
		users = append(users, u)
	}

	return users
}
