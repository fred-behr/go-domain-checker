package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fred-behr/go-domain-checker/models"
)

func CheckByAPI(dC models.DomainChecker, d models.Domain) bool {
	// Setting up request with key and query
	req, _ := http.NewRequest("GET", dC.GetQuery(d), nil)
	req.Header.Add("x-api-key", dC.Key)

	// Calling API
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Domain API failed")
		return false
	}
	defer res.Body.Close()

	// Checking response
	body, _ := io.ReadAll(res.Body)
	var result map[string]interface{} // JSON ends up here
	json.Unmarshal(body, &result)
	available, ok := result["available"].(bool)
	if ok {
		return available
	}
	return false
}

func CheckDomain(d models.Domain, dC models.DomainChecker, ch chan models.DomainResult) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest("GET", d.GetURL(), nil)
	_, err := client.Do(req)
	if err != nil {
		ch <- models.DomainResult{
			Domain:      d,
			IsAvailable: CheckByAPI(dC, d),
			IsReachable: false,
			IsForSale:   false, //todo: Check
			Error:       err,
		}
		return
	}
	ch <- models.DomainResult{
		Domain:      d,
		IsAvailable: false,
		IsReachable: true,
		IsForSale:   false, //todo: Check
		Error:       nil,
	}
}
