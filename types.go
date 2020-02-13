package golocator

//go:generate msgp
type Country struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	Iso3      string              `json:"iso3"`
	Iso2      string              `json:"iso2"`
	PhoneCode string              `json:"phone_code"`
	Capital   string              `json:"capital"`
	Currency  string              `json:"currency"`
	States    map[string][]string `json:"states"`
}

type GoLocator struct {
	Countires []Country
}
