package controller

import (
	"log"
	"net/http"
)

func NewRouter(server *http.Server) {
	mux := http.NewServeMux()

	mux.HandleFunc("/taxid/", server.taxHandler)

	log.Fatal(http.ListenAndServe("localhost:"+"6000", mux))
}
