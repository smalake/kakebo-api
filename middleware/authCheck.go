package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/smalake/kakebo-api/utils/firebase"
	"github.com/smalake/kakebo-api/utils/logging"
)

type MyKey string

// FirebaseのJWTを検証する
func AuthCheck(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// クライアントから送られてきた JWT 取得
		// authHeader := r.Header.Get("Authorization")
		// idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// CookieからJWTを取得
		cookie, err := r.Cookie("kakebo")
		if err != nil {
			logging.WriteErrorLog(err.Error(), true)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "トークンが取得できません")
			return
		}

		idToken := cookie.Value

		// JWT の検証
		token, err := firebase.Auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			logging.WriteErrorLog(err.Error(), true)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, err)
			return
		}

		// http.RequestのContextにUIDを設定
		ctx := context.WithValue(r.Context(), MyKey("uid"), token.UID)
		// コンテキストを設定したhttp.Requestを次のハンドラに渡す
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
