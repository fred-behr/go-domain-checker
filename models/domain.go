package models

type Domain struct {
	Name           string
	TopLevelDomain string
}

func (d Domain) GetURL() string {
	return "http://www." + d.Name + d.TopLevelDomain
}

func (d Domain) GetDomain() string {
	return d.Name + d.TopLevelDomain
}
