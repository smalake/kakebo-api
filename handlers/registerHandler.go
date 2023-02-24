package handlers

import (
	"encoding/json"
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
		return
	}
	// ユーザの新規登録
	err2 := user.RegisterUser()
	if err2 != nil {
		logging.WriteErrorLog(err2.Error(), true)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// JWTの生成
	tokenString, err := jwt.CreateJWT(user.Email)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// JWTをクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}
