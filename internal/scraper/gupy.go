package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/1-AkM-0/empreGo-web/internal/models"
)

type gupyJobs struct {
	Data []struct {
		Title string `json:"name"`
		Link  string `json:"jobUrl"`
	} `json:"data"`
}

func SearchGupy(jobChannel chan models.Job) error {
	rawUrl := "https://employability-portal.gupy.io/api/v1/jobs?jobName=est%C3%A1gio&limit=10&offset=0&workplaceType=remote"
	method := "GET"

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, rawUrl, nil)
	if err != nil {
		return fmt.Errorf("erro na tentativa de fazer o wrapper do request: %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erro ao fazer o request: %v", err)
	}
	defer res.Body.Close()
	var gupyResponse gupyJobs
	err = json.NewDecoder(res.Body).Decode(&gupyResponse)
	if err != nil {
		return fmt.Errorf("erro ao tentar decodar o json: %v", err)
	}

	for _, result := range gupyResponse.Data {
		if !(isTechInternship(result.Title)) {
			continue
		}
		jobToInsert := models.Job{
			Title:  result.Title,
			Link:   result.Link,
			Source: "gupy",
			Type:   findJobType(result.Title),
		}
		jobChannel <- jobToInsert
	}
	return nil
}
