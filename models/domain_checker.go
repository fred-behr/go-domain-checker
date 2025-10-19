package models

import (
	"os"

	"github.com/joho/godotenv"
)

type DomainChecker struct {
	BaseURL string
	Key     string
}

func NewDomainChecker() DomainChecker {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	return DomainChecker{
		BaseURL: "https://api.api-ninjas.com/v1/domain?domain=",
		Key:     apiKey,
	}
}

func (d DomainChecker) GetQuery(Domain Domain) string {
	return d.BaseURL + Domain.GetDomain()
}
