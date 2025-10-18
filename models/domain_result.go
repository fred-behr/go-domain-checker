package models

type DomainResult struct {
	Domain      Domain
	IsReachable bool
	IsForSale   bool
	Error       error
}
