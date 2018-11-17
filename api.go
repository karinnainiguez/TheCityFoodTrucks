package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/SebastiaanKlippert/go-soda"
)

type truck struct {
	Applicant    string `json:"applicant"`
	Coldtruck    string `json:"coldtruck"`
	Dayofweekstr string `json:"dayofweekstr"`
	End24        string `json:"end24"`
	Endtime      string `json:"endtime"`
	Location     string `json:"location"`
	Locationdesc string `json:"locationdesc,omitempty"`
	Start24      string `json:"start24"`
	Starttime    string `json:"starttime"`
	Optionaltext string `json:"optionaltext,omitempty"`
}

func getQuery() []truck {
	sodareq := soda.NewGetRequest("https://data.sfgov.org/resource/bbb8-hzi6", "")

	now := time.Now()

	sodareq.Format = "json"
	sodareq.Query.Select = []string{
		"applicant",
		"starttime",
		"start24",
		"end24",
		"endtime",
		"dayofweekstr",
		"location",
		"optionaltext",
		"coldtruck",
		"locationdesc",
	}
	sodareq.Filters["dayofweekstr"] = now.Weekday().String()

	resp, err := sodareq.Get()
	handle(err)

	defer resp.Body.Close()

	trucks := make([]truck, 0)
	respBytes, err := ioutil.ReadAll(resp.Body)
	handle(err)
	json.Unmarshal(respBytes, &trucks)

	filtered, err := filterTrucksByTime(now.Format("15:04"), trucks)

	return filtered
}

func filterTrucksByTime(timeNow string, totalColl []truck) ([]truck, error) {

	newTrucks := make([]truck, 0)

	hourNow, err := time.Parse("15:04", timeNow)
	handle(err)

	for _, tr := range totalColl {

		startString := tr.Start24
		if startString == "24:00" {
			startString = "00:00"
		}

		truckStart, err := time.Parse("15:04", startString)
		handle(err)

		endString := tr.End24
		if endString == "24:00" {
			endString = "00:00"
		}

		truckEnd, err := time.Parse("15:04", endString)
		handle(err)

		if truckStart.Before(hourNow) && truckEnd.After(hourNow) {
			newTrucks = append(newTrucks, tr)
		}
	}

	return newTrucks, nil
}
