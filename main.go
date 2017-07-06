package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	code := os.Args[1]

	resp, err := http.PostForm("http://www2.correios.com.br/sistemas/rastreamento/newprint.cfm",
		url.Values{"objetos": {code}})

	if err != nil {
		log.Fatal("Cannot open url: ", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	events := []string{}

	doc.Find(".listEvent tr").Each(func(i int, s *goquery.Selection) {
		text := s.Find(".sroLbEvent").Text()
		info := s.Find(".sroDtEvent").Text()
		r, _ := regexp.Compile("\\s+")
		text = r.ReplaceAllString(text, " ")
		info = r.ReplaceAllString(info, " ")
		events = append(events, info+" - "+text)
	})

	for i := len(events) - 1; i >= 0; i-- {
		fmt.Println(events[i])
	}

}
