package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kalunik/siteWrapper/internal/models"
)

func main() {
	comp := &models.Company{}
	//var t map[string]interface{}

	resp, err := http.Get("https://www.rusprofile.ru/ajax.php?&query=3327848813&action=search")
	if err != nil {
		log.Println(err)
		return
	}
	if resp.StatusCode > 400 {
		log.Println("Get failed", resp.Status)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &comp)
	if err != nil {
		log.Println("Unmarshal failed", err)
		return
	}

	//fmt.Println("Data", r.UL.Inn, r.UL.Name, r.UL.RawName, r.UL.CeoName)
	//fmt.Println(t)

	fmt.Println(comp.Ul)
	//out, _ := json.Marshal(r)
	//fmt.Println(string(out))
}
