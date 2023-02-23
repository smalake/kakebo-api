package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/jwt"
	"github.com/smalake/kakebo-api/utils/logging"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからユーザー名とパスワードを取得
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// パスワードを検証し、認証に成功すればJWTを生成する
	pass, err := model.GetPassword(user.Email)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		return
	}
	if user.Password == pass {
		// JWTの生成
		tokenString, err := jwt.CreateJWT(user.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// JWTをクライアントに返す
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
