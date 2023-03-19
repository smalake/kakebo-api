package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/smalake/kakebo-api/model"
)

func EventHandler(w http.ResponseWriter, r *http.Request) {
	var events model.Events
	var event model.Event
	// コンテキストからUIDを取得
	uid := r.Context().Value("uid").(string)

	// リクエストの種類に応じて処理を実行
	switch r.Method {
	case http.MethodGet: // GETリクエストの処理
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
		return

	case http.MethodPost: // POSTリクエストの処理
		// リクエストボディから登録内容を取得
		err := json.NewDecoder(r.Body).Decode(&events)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		err = events.CreateEvent(uid)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPut: // PUTリクエストの処理
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
		return
	case http.MethodDelete: // DELETEリクエストの処理
		// リクエストボディから削除内容を取得
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		err = event.DeleteEvent()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
