package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/jwt"
	"github.com/smalake/kakebo-api/utils/logging"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからユーザー名とパスワードを取得
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	// パスワードを検証し、認証に成功すればJWTを生成する
	pass, err := user.GetPassword()
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		fmt.Fprintln(w, err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password)) //ハッシュ化されたパスワードを比較
	// パスワードが一致しなかった場合
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, err)
		return
	}

	// JWTの生成
	tokenString, err := jwt.CreateJWT(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	// JWTをクライアントに返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}