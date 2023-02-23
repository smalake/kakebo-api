package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/smalake/kakebo-api/model"
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
	if user.Password == model.GetPassword(user.Email) {
		// JWTの生成
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 168).Unix(), //有効期限1週間
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
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
