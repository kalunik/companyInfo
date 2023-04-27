package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/kalunik/companyInfo/internal/parse"
	"log"
	"net/http"
)

func TaxHandler(w http.ResponseWriter, r *http.Request) {
	taxId := chi.URLParam(r, "taxId")

	info := parse.ScrapeRusprofile(taxId)
	fmt.Println(info)
	data, err := json.Marshal(info)
	if err != nil {
		log.Fatalln("Error while marshal response", err.Error())
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalln("Failede to write response", err.Error())
	}

}
