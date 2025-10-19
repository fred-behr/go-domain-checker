package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fred-behr/go-domain-checker/models"
	"github.com/fred-behr/go-domain-checker/utils"
)

func main() {
	domains, err := getDomains()
	if err != nil {
		fmt.Println("Failed to make list of domain names:", err)
		os.Exit(1)
	}

	// Get results & print out to stdout
	for _, domain := range domains {
		result := checkDomain(domain)
		result.PrintResult()
	}
}

func getDomains() ([]models.Domain, error) {
	// Get domain names from input file
	lines, err := utils.ReadFile("./input/domain_names.txt")
	if err != nil {
		fmt.Println("Failed to get domains from input file:", err)
		return nil, err
	}

	// Get top level domain names from input file
	topLevelDomains, err := getTopLevelDomains()
	if err != nil {
		return nil, err
	}

	// Create list of full domain names
	domains := []models.Domain{}
	for _, line := range lines {
		domain := strings.TrimSpace(strings.ToLower(line))

		// Replacing empty spaces with "-", so it is a valid URL
		domain = strings.ReplaceAll(domain, " ", "-")

		if domain == "" {
			continue
		}
		for _, topLevelDomain := range topLevelDomains {
			domains = append(domains, models.Domain{
				Name:           domain,
				TopLevelDomain: topLevelDomain,
			})
		}
	}
	return domains, nil
}

func getTopLevelDomains() ([]string, error) {
	lines, err := utils.ReadFile("./input/top_level_domains.txt")
	if err != nil {
		fmt.Println("Failed to get top level domains from input file:", err)
		return nil, err
	}
	topLevelDomains := []string{}
	for _, line := range lines {
		topLevelDomains = append(topLevelDomains, strings.TrimSpace(line))
	}
	return topLevelDomains, nil
}

func checkDomain(domain models.Domain) models.DomainResult {
	_, err := http.Get(domain.GetURL())
	if err != nil {
		return models.DomainResult{
			Domain:      domain,
			IsReachable: false,
			IsForSale:   false, //todo: Check
			Error:       err,
		}
	}
	return models.DomainResult{
		Domain:      domain,
		IsReachable: true,
		IsForSale:   false, //todo: Check
		Error:       nil,
	}
}
