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

// グループを新規作成
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	// コンテキストからUIDを取得
	uid := r.Context().Value("uid").(string)

	var setupGroup setup.Group

	err := setupGroup.CreateGroup(uid)
	if err != nil {
		logging.WriteErrorLog(err.Error(), true)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)

}
