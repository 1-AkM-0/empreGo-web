package scraper

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/1-AkM-0/empreGo-web/internal/models"
	"github.com/PuerkitoBio/goquery"
)

func SearchLinkedin(jobChannel chan models.Job) error {
	rawUrl := "https://www.linkedin.com/jobs/search?keywords=%22est%C3%A1gio%22%20OR%20%22estagi%C3%A1rio%22&location=Brasil&geoId=106057199&f_TPR=r86400&f_WT=2&position=1&pageNum=0&currentJobId=4373363527"
	method := "GET"

	client := http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, rawUrl, nil)
	if err != nil {
		return fmt.Errorf("linkedin NewRequest: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, Like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	res, err := client.Do(req)
	if res.StatusCode != 200 {
		return fmt.Errorf("linkedin status: %w", err)
	}
	if err != nil {
		return fmt.Errorf("linkedin clientDo: %w", err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("linkedin html parser: %w", err)
	}

	doc.Find("ul.jobs-search__results-list > li").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h3.base-search-card__title").Text())
		link, exists := (s.Find("a.base-card__full-link").Attr("href"))
		u, err := url.Parse(link)
		if err != nil {
			return
		}
		u.RawQuery = ""
		link = u.String()

		if title != "" && exists {
			job := models.Job{
				Title:  title,
				Link:   link,
				Source: "linkedin",
			}
			jobChannel <- job
		}
	})
	return nil
}
