package golocator

import (
	"errors"
	"os"
	"strings"

	"github.com/tinylib/msgp/msgp"
)

type GoLocatorService struct {
	DB *GoLocator
}

func NewGoLocatorService(path string) GoLocatorService {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var db GoLocator

	err = msgp.ReadFile(&db, file)

	if err != nil {
		panic(err)
	}

	return GoLocatorService{
		DB: &db,
	}
}

func (g *GoLocatorService) FindCountryByIso2(countryCode string) (Country, error) {
	countryCode = strings.ToUpper(countryCode)
	for _, country := range g.DB.Countires {
		if country.Iso2 == countryCode {
			return country, nil
		}
	}
	return Country{}, errors.New("country not found")
}

func (g *GoLocatorService) FindCountryByIso3(countryCode string) (Country, error) {
	countryCode = strings.ToUpper(countryCode)
	for _, country := range g.DB.Countires {
		if country.Iso3 == countryCode {
			return country, nil
		}
	}
	return Country{}, errors.New("country not found")
}

func (g *GoLocatorService) FindCities(country Country) []string {
	var cities []string
	for _, state := range country.States {
		for _, city := range state {
			cities = append(cities, city)
		}

	}
	return cities
}
