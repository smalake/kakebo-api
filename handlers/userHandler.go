package handlers

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/sessions"
// 	"github.com/smalake/kakebo-api/model"
// )

// // セッション用のストア
// var store = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_KEY")))

// func UserHandler(w http.ResponseWriter, _ *http.Request) {
// 	users := model.GetUser()
// 	fmt.Fprintln(w, users)
// }

// func UserDataHandler(w http.ResponseWriter, r *http.Request) {
// 	// セッションからユーザIDを取得
// 	session, _ := store.Get(r, os.Getenv("SESSION_NAME"))
// 	id := session.Values["id"].(int)

// 	data := model.GetUserData(id)
// 	fmt.Fprintln(w, data.Name)

// }
