package handler

import (
	"fmt"
	"net/http"
	"os"

	utils "social-network/pkg"
)

func (Handler *Handler) ServeFilesHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = r.URL.Path[1:]
	_, err := os.ReadFile(r.URL.Path)
	if err != nil {
		fmt.Println(err)
		fmt.Println(r.URL.Path)
		utils.WriteJson(w, http.StatusInternalServerError, "StatusInternalServerError")
		return
	}

	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
}
