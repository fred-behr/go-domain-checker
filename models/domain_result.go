package models

import "fmt"

type DomainResult struct {
	Domain      Domain
	IsAvailable bool
	IsReachable bool
	IsForSale   bool
	Error       error
}

func (d DomainResult) PrintResult() {
	if d.IsAvailable {
		fmt.Printf("POSITIVE: %s is available.\n", d.Domain.GetDomain())
		return
	}
	if d.IsForSale {
		fmt.Printf("FOR_SALE: %s is for sale.\n", d.Domain.GetDomain())
		return
	}
	if d.IsReachable {
		fmt.Printf("NEGATIVE: %s is NOT for sale.\n", d.Domain.GetDomain())
		return
	} else {
		fmt.Printf("NOTKNOWN: %s might be available.\n", d.Domain.GetDomain())
		return
	}
}
