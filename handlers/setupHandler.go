package handlers

import (
	"fmt"
	"net/http"

	"github.com/smalake/kakebo-api/model/setup"
	"github.com/smalake/kakebo-api/utils/logging"
)

// セットアップしてあるユーザかチェックするためグループIDを取得
func CheckSetup(w http.ResponseWriter, r *http.Request) {
	// コンテキストからUIDを取得
	uid := r.Context().Value("uid").(string)

	var setup setup.Group

	groupData, err := setup.GetGroupID(uid)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	// 取得したイベントをフロント側へと渡す
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(groupData))
}
