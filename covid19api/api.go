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

func fetchInto(v interface{}, endpoint string) error {
	body, err := fetch(endpoint)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &v)
}

func GetSummary() (Summary, error) {
	var s Summary

	err := fetchInto(&s, "summary")

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

	err := fetchInto(&cs, endpoint)

	return cs, err
}

func GetLiveByCountryAndStatus(country, status string) ([]ByCountry, error) {
	var cs []ByCountry

	endpoint := fmt.Sprintf("live/country/%s/status/%s", country, status)

	err := fetchInto(&cs, endpoint)

	return cs, err
}

func GetByCountry(country, status string) ([]ByCountry, error) {
	var cs []ByCountry

	endpoint := fmt.Sprintf("country/%s/status/%s", country, status)

	err := fetchInto(&cs, endpoint)

	return cs, err
}

func GetByCountryLive(country, status string) ([]ByCountry, error) {
	var cs []ByCountry

	endpoint := fmt.Sprintf("country/%s/status/%s/live", country, status)

	err := fetchInto(&cs, endpoint)

	return cs, err
}

func GetByCountryTotal(country, status string) ([]ByCountry, error) {
	var cs []ByCountry

	endpoint := fmt.Sprintf("total/country/%s/status/%s", country, status)

	err := fetchInto(&cs, endpoint)

	return cs, err
}
