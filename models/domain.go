package models

type Domain struct {
	Name           string
	TopLevelDomain string
}

func (d Domain) GetURL() string {
	return "http://www." + d.Name + d.TopLevelDomain
}
