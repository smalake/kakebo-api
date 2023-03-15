package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smalake/kakebo-api/model"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからUIDを取得
	var user model.User
	fmt.Println(user)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	// ログイン処理
	err = user.LoginUser()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		// fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
