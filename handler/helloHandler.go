package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
