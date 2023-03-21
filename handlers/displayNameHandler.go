package handlers

import (
	"fmt"
	"net/http"
)

func DisplayNameHandler(w http.ResponseWriter, r *http.Request) {
	// var displayName model.DisplayName
	// コンテキストからUIDを取得
	// uid := r.Context().Value("uid").(string)

	switch r.Method {
	case http.MethodGet:
		q := r.URL.Query()
		fmt.Println(q)
		return
	case http.MethodPut:

	}
}
