package internal

import (
	"net/http"
)

func NewRouter(server *http.Server) {
	mux := http.NewServeMux()

	mux.HandleFunc("/taxid/", taxHandler)

}