package golocator

import (
	"testing"
)

func TestFindCountryByIso2(t *testing.T) {
	locator := NewGoLocatorService("golocator.db")

	country, _ := locator.FindCountryByIso2("US")

	if country.Name != "United States" {
		t.Error("find by iso2 error")
	}

}

func TestFindCountryByIso3(t *testing.T) {
	locator := NewGoLocatorService("golocator.db")

	country, _ := locator.FindCountryByIso3("USA")

	if country.Name != "United States" {
		t.Error("find by iso3 error")
	}

}
