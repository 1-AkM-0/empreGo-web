package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/1-AkM-0/empreGo-web/internal/models"
)

type gupyJobs struct {
	Data []struct {
		Title   string `json:"name"`
		Link    string `json:"jobUrl"`
		Company string `json:"careerPageName"`
	} `json:"data"`
}

func SearchGupy(jobChannel chan models.Job) error {
	var wg sync.WaitGroup

	gupyChannel := make(chan *http.Response, 2)

	rawUrls := []string{"https://employability-portal.gupy.io/api/v1/jobs?jobName=est%C3%A1gio&limit=10&offset=0&workplaceType=remote", "https://employability-portal.gupy.io/api/v1/jobs?jobName=estagi%C3%A1rio&limit=10&offset=0&workplaceType=remote"}

	for _, url := range rawUrls {
		wg.Go(func() {
			res, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			gupyChannel <- res
		})
	}

	go func() {
		wg.Wait()
		close(gupyChannel)
	}()

	for response := range gupyChannel {
		gupyResponse := &gupyJobs{}

		err := json.NewDecoder(response.Body).Decode(&gupyResponse)
		if err != nil {
			return fmt.Errorf("erro ao tentar decodar o json: %v", err)
		}

		err = response.Body.Close()
		if err != nil {
			return fmt.Errorf("erro ao fechar o body da requisição Gupy")
		}
		for _, result := range gupyResponse.Data {
			if !(isTechInternship(result.Title)) {
				continue
			}
			jobToInsert := models.Job{
				Title:   result.Title,
				Link:    result.Link,
				Source:  "gupy",
				Type:    findJobType(result.Title),
				Company: result.Company,
			}
			jobChannel <- jobToInsert
		}
	}
	return nil
}
