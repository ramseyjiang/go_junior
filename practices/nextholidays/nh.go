package nh

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type holiday struct {
	Date        string      `json:"date"`
	LocalName   string      `json:"localName"`
	Name        string      `json:"name"`
	CountryCode string      `json:"countryCode"`
	Counties    interface{} `json:"counties"`
	LaunchYear  interface{} `json:"launchYear"`
	Type        string      `json:"type"`
	Fixed       bool        `json:"fixed"`
	Global      bool        `json:"global"`
}

var nzHolidays []*holiday

const url = "https://date.nager.at/api/v2/publicholidays/2022/NZ"

func GetHolidays() (map[string]*holiday, error) {
	resp, _ := http.Get(url)
	payload, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(payload, &nzHolidays)
	if err != nil {
		return nil, err
	}

	nHolidays := make(map[string]*holiday)
	currentDateStr := time.Now().Format("20060102")

	for _, h := range nzHolidays {
		strKey := strings.Replace(h.Date, "-", "", 2)
		if currentDateStr <= strKey && h.Global {
			nHolidays[strKey] = h
		}
	}

	return nHolidays, nil
}