package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/fred-behr/go-domain-checker/models"
	"github.com/fred-behr/go-domain-checker/utils"
)

func main() {
	domains, err := getDomains()
	if err != nil {
		fmt.Println("Failed to make list of domain names:", err)
		os.Exit(1)
	}

	// Init DomainChecker
	domainChecker := models.NewDomainChecker()

	ch := make(chan models.DomainResult)
	sem := make(chan struct{}, 10) // limit concurrency
	var wg sync.WaitGroup

	for _, domain := range domains {
		wg.Add(1)
		go func(d models.Domain) {
			defer wg.Done()
			sem <- struct{}{} // acquire
			utils.CheckDomain(d, domainChecker, ch)
			<-sem // release
		}(domain)
	}

	// Close channel when all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Print results
	for result := range ch {
		go func(result models.DomainResult) {
			result.PrintResult()
		}(result)
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
