package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smalake/kakebo-api/model/events"
)

// 対象ユーザのイベント一覧を取得
func GetEvents(w http.ResponseWriter, r *http.Request) {
	var event events.GetEvent
	// コンテキストからUIDを取得
	uid := r.Context().Value("uid").(string)

	// 一覧を取得
	eventList, err := event.GetEvents(uid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	// 取得したイベントをフロント側へと渡す
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(eventList))
}

// イベントを新規作成
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event events.CreateEvent
	// コンテキストからUIDを取得
	uid := r.Context().Value("uid").(string)

	// リクエストボディから登録内容を取得
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	err = event.InsertEvent(uid)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// イベントを更新
func EditEvent(w http.ResponseWriter, r *http.Request) {
	var event events.UpdateEvent

	// リクエストボディから更新内容を取得
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	err = event.UpdateEvent()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var event events.DeleteEvent

	// リクエストボディから削除内容を取得
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	err = event.DeleteOneEvent()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
