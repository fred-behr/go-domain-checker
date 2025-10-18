package models

import "fmt"

type DomainResult struct {
	Domain      Domain
	IsReachable bool
	IsForSale   bool
	Error       error
}

func (d DomainResult) PrintResult() {
	if d.IsForSale {
		fmt.Printf("POSITIVE: %s is for sale.\n", d.Domain.GetURL())
		return
	}
	if d.IsReachable {
		fmt.Printf("NEGATIVE: %s is NOT for sale.\n", d.Domain.GetURL())
		return
	} else {
		fmt.Printf("POSITIVE: %s might be available.\n", d.Domain.GetURL())
		return
	}
}
