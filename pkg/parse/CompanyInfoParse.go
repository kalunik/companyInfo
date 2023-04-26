package parse

import (
	"fmt"
	"log"
	"net/url"
)

type Company struct {
	TaxId   string
	KPP     string
	Name    string
	CeoName string
}

func ScrapeRusprofile(taxId string) *Company {
	targetURL, err := url.Parse("https://www.rusprofile.ru/search?query=")
	if err != nil {
		log.Fatalln("URL parse failed ", err.Error())
	}
	updateQuery(targetURL, taxId)

	s := new(Scrapper)
	data := s.scrapeSite(targetURL.String()) //3327848813
	fmt.Println(data)

	company := new(Company)
	company.Unmarshal(data)
	return company
}

func updateQuery(url *url.URL, value string) {
	q := url.Query()
	q.Set("query", value)
	url.RawQuery = q.Encode()
}

func (info *Company) Unmarshal(data []string) {
	for i, field := range data {
		switch i {
		case 0:
			info.TaxId = field
		case 1:
			info.KPP = field
		case 2:
			info.Name = field
		case 3:
			info.CeoName = field
		case 4:
			break
		}
	}
}
