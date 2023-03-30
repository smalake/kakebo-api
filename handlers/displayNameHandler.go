package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smalake/kakebo-api/model"
	"github.com/smalake/kakebo-api/utils/logging"
)

// ユーザの表示名を取得する
func GetName(w http.ResponseWriter, r *http.Request) {
	var displayName model.DisplayName
	// コンテキストからUIDを取得
	uid := r.Context().Value("uid").(string)

	name, err := displayName.GetDisplayName(uid)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	// 取得したイベントをフロント側へと渡す
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(name))
}

// ユーザの表示名を変更する
func EditName(w http.ResponseWriter, r *http.Request) {
	var displayName model.DisplayName
	// コンテキストからUIDを取得
	uid := r.Context().Value("uid").(string)

	// リクエストボディから更新内容を取得
	err := json.NewDecoder(r.Body).Decode(&displayName)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	err = displayName.UpdateDisplayName(uid)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
