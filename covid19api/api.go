package covid19api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CountrySummary struct {
	Country        string
	Slug           string
	TotalConfirmed int
	TotalDeaths    int
	TotalRecovered int
	NewConfirmed   int
	NewDeaths      int
	NewRecovered   int
}

type Summary struct {
	Countries []CountrySummary `json:"Countries"`
	Date      time.Time        `json:"Date"`
}

func fetch(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("https://api.covid19api.com/%s", endpoint)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func GetSummary() (Summary, error) {
	var s Summary

	body, err := fetch("summary")
	if err != nil {
		return s, nil
	}

	err = json.Unmarshal(body, &s)

	return s, err
}

type ByCountry struct {
	Country   string
	Province  string
	Latitude  float32 `json:"Lat"`
	Longitude float32 `json:"Lon"`
	Date      time.Time
	Cases     int
	Status    string
}

func GetLiveByCountryAndStatusAfterDate(country, status string, date time.Time) ([]ByCountry, error) {
	var cs []ByCountry

	endpoint := fmt.Sprintf("live/country/%s/status/%s/date/%s", country, status, date.Format(time.RFC3339))

	body, err := fetch(endpoint)
	if err != nil {
		return nil, nil
	}

	err = json.Unmarshal(body, &cs)

	return cs, err
}