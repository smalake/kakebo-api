package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/smalake/kakebo-api/model"
)

// セッション用のストア
var store = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_KEY")))

func AuthCheck(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// トークンの形式を検証
		bearerToken := strings.Split(tokenHeader, " ")
		if len(bearerToken) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// JWTの検証
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// JWTからメールアドレスを抽出
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			// 抽出したメールアドレスからユーザIDを取得
			id := model.GetUserId(email)
			// セッションを開始
			session, _ := store.Get(r, os.Getenv("SESSION_NAME"))
			session.Values["id"] = id //ユーザIDをセッションに保存
			session.Save(r, w)
		}
		next.ServeHTTP(w, r)
	})
}
