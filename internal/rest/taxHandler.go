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

	if len(taxId) < 10 || len(taxId) > 12 {
		w.WriteHeader(http.StatusBadRequest)
		err := requestError{wrongLenTaxId, nil}
		w.Write(err.ErrorJson())
		return
	}

	info := parse.ScrapeRusprofile(taxId)
	fmt.Println(info)
	data, err := json.Marshal(info)
	if err != nil {
		log.Fatalln("Error while marshal response", err.Error())
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalln("Failed to write response", err.Error())
	}

}
