package parse

import (
	"github.com/gocolly/colly/v2"
	"log"
)

type Scrapper struct {
	collector *colly.Collector
}

func (c *Scrapper) scrapeSite(targetURL string) []string {
	c.collector = colly.NewCollector()

	tagsList := []string{
		"span[id=clip_inn]",
		"span[id=clip_kpp]",
		"div[class=company-name]",
		"span[class=company-info__text] > a > span",
	}
	collectedData := make([]string, 0, len(tagsList))
	for _, tag := range tagsList {
		c.collector.OnHTML(tag, func(e *colly.HTMLElement) {
			collectedData = append(collectedData, e.Text)
		})
	}
	c.collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})
	err := c.collector.Visit(targetURL)
	if err != nil {
		log.Fatalln("Collecting failed while scrapping site", err.Error())
	}
	return collectedData
}
