package scraper

import (
	"regexp"
	"slices"
	"strings"
)

func isTechInternship(jobTitle string) bool {

	blacklist := []string{"marketing", "negócios", "growth", "contabilidade", "civil", "customer", "vendas", "videomaker", "suporte operacional", "suporte comercial", "cro", "comercial", "monetização", "cx", "desenvolvimento comercial"}

	jobTitle = strings.ToLower(jobTitle)
	re := regexp.MustCompile(`[\- :,()]`)
	cleanString := re.ReplaceAllString(jobTitle, " ")
	words := strings.Fields(cleanString)
	keywordsIntership := []string{"estágio", "estagio", "estagiário", "estagiario", "estagiária", "estagiaria", "intern", "internship"}
	isInternship := false
	isTech := false

	for _, keyword := range words {
		if slices.Contains(blacklist, keyword) {
			return false
		}
	}

	for _, keyword := range words {
		if slices.Contains(keywordsIntership, keyword) {
			isInternship = true
			break
		}
	}

	if !isInternship {
		return false
	}

	keywords := []string{
		"dados", "ti", "backend", "fullstack", "full", "suporte", "devops", "desenvolvimento", "software",
		"frontend", "front", "front end", "mobile", "android", "ios", "web", "back end", "computação",
		"desenvolvedor", "developer", "programador", "engenheiro", "engineer", "arquiteto", "sistemas", "programação",
		"data", "machine learning", "ia", "ai", "analytics", "dev",
		"cloud", "qa", "testes", "teste",
	}
	for _, keyword := range words {
		isTech = slices.Contains(keywords, keyword)
		if isTech {
			break
		}
	}

	return isTech
}

func findJobType(jobTitle string) string {
	jobTitle = strings.ToLower(jobTitle)
	re := regexp.MustCompile(`[\- :,()]`)
	cleanString := re.ReplaceAllString(jobTitle, " ")
	for words := range strings.FieldsSeq(cleanString) {
		switch words {
		case "backend", "back":
			return "backend"
		case "fullstack", "full":
			return "fullstack"
		case "frontend", "front":
			return "frontend"
		}
	}
	return "geral"
}
