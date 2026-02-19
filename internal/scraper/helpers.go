package scraper

import (
	"regexp"
	"slices"
	"strings"
)

func isTechInternship(jobTitle string) bool {
	jobTitle = strings.ToLower(jobTitle)
	re := regexp.MustCompile(`[\- :,()]`)
	cleanString := re.ReplaceAllString(jobTitle, " ")
	words := strings.Fields(cleanString)
	keywordsIntership := "estágio estagio estagiário estagiario estagiária estagiaria intern internship"
	isInternship := false
	isTech := false

	for _, keyword := range words {
		if strings.Contains(keywordsIntership, keyword) {
			isInternship = true
			break
		}
	}

	if !isInternship {
		return false
	}

	keywords := []string{"dados", "ti", "backend", "fullstack", "full-stack", "suporte", "devops", "desenvolvimento", "software"}
	for _, keyword := range words {
		isTech = slices.Contains(keywords, keyword)
		if isTech {
			break
		}
	}
	return isTech
}
