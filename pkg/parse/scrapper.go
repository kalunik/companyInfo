package parse

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type CompanyInfo struct {
	TaxId   string
	KPP     string
	Name    string
	CeoName string
}

func ScrapeRusprofile(taxId string) {
	scrape("https://www.rusprofile.ru/search?query=" + taxId) //3327848813
}

func scrape(url string) *CompanyInfo {
	c := colly.NewCollector()

	company := new(CompanyInfo)
	// Find and visit all links
	c.OnHTML("div[class=company-name],,", func(e *colly.HTMLElement) {
		company.Name = e.Text
	})

	c.OnHTML("span[id=clip_inn]", func(e *colly.HTMLElement) {
		company.TaxId = e.Text
	})
	c.OnHTML("span[id=clip_kpp]", func(e *colly.HTMLElement) {
		company.KPP = e.Text
	})
	c.OnHTML("span[class=company-info__text] > a > span", func(e *colly.HTMLElement) {
		company.CeoName = e.Text
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)
	return company
}
