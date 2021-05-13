// リクエストを変換し、modelのロジックを呼び出し、viewに変換する
package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"tutorial/mvc/pkg/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	sample := model.Sample{Message: "modelの導入できた"}
	json.NewEncoder(w).Encode(sample)
	fmt.Fprintf(w, "テストやで！")
	fmt.Println(sample)
	fmt.Println("終了")
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Postの処理を書いていくよ")
	db := model.GormConnect()
	user := model.User{Name: "shinya"}
	db.Create(&user)
	// defer db.Close()
	fmt.Println(user.Name)
	fmt.Fprintf(w, "いけるかな...")
}