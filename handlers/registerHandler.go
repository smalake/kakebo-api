package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/jwt"
	"github.com/smalake/kakebo-api/utils/logging"
)

// ユーザの新規登録
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからメールアドレス・ユーザ名・パスワードを取得
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	// ユーザの新規登録
	err = user.RegisterUser()
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	// JWTの生成
	tokenString, err := jwt.CreateJWT(user.Email)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	// JWTをクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}