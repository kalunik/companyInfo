package main

import (
	"github.com/kalunik/companyInfo"
	"github.com/kalunik/companyInfo/internal"
	"log"
)

func main() {

	mux := internal.NewRouter()

	server := new(companyInfo.Server)
	err := server.Run("8080", mux)
	if err != nil {
		log.Fatalf("error occured while running http server: %s\n", err.Error())
	}



	//
	//server := &NewCompanyInfoServer()

	//controller.NewRouter(server)

	/*//comp := &models.Company{}
	var t map[string]interface{}
	//https://www.rusprofile.ru/search?query=3327848813
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

	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Println("Unmarshal failed", err)
		return
	}

	//fmt.Println("Data", r.UL.Inn, r.UL.Name, r.UL.RawName, r.UL.CeoName)
	//fmt.Println(t)

	fmt.Println("_________!", string(body))
	fmt.Println(t)
	//out, _ := json.Marshal(r)
	//fmt.Println(string(out))*/
}
