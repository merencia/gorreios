package crawler

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/apex/log"
	"github.com/merencia/gorreios/model"
)

func GetPackage(code string) []model.Event {
	resp, err := http.PostForm("http://www2.correios.com.br/sistemas/rastreamento/newprint.cfm",
		url.Values{"objetos": {code}})

	if err != nil {
		log.WithError(err).Fatal("cannot open url ")
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.WithError(err).Info("cannot parse html")
	}

	events := []model.Event{}
	r, _ := regexp.Compile("\\s+")
	dateReg := regexp.MustCompile("\\d{2}\\/\\d{2}\\/\\d{4}\\s\\d{2}:\\d{2}")
	locationReg := regexp.MustCompile("[a-zA-Z].+\\s{0,}\\/.{0,}[A-Z]+")

	doc.Find(".listEvent tr").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Find(".sroLbEvent").Text())
		info := s.Find(".sroDtEvent").Text()
		text = r.ReplaceAllString(text, " ")
		info = r.ReplaceAllString(info, " ")

		dateText := dateReg.FindStringSubmatch(info)[0]
		date, _ := time.Parse("02/01/2006 15:04", dateText)
		location := locationReg.FindStringSubmatch(info)[0]

		events = append(events, model.Event{
			Date:     date,
			Location: location,
			Event:    text,
		})
	})

	return events
}
