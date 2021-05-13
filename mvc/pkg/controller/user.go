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
	
}