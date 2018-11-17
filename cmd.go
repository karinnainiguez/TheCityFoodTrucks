package main

import (
	"fmt"
	"os"

	"github.com/kubicorn/kubicorn/pkg/logger"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
)

const RESULTSPERPAGE = 10

func root(allTrucks []truck, pageNum int) {
	// Display Food Trucks Currently Open
	showTrucks(allTrucks, pageNum)

	// display options
	var prompt promptui.Select

	if pageNum == 0 && len(allTrucks)/RESULTSPERPAGE < pageNum {
		prompt = promptui.Select{
			Label: fmt.Sprintf("Page No. %v  Options", pageNum+1),
			Items: []string{"Exit"},
		}
	} else if pageNum == 0 {
		prompt = promptui.Select{
			Label: fmt.Sprintf("Page No. %v  Options", pageNum+1),
			Items: []string{"Next", "Exit"},
		}
	} else if len(allTrucks)/RESULTSPERPAGE <= pageNum {
		prompt = promptui.Select{
			Label: fmt.Sprintf("Page No. %v  Options", pageNum+1),
			Items: []string{"Previous", "Exit"},
		}
	} else {
		prompt = promptui.Select{
			Label: fmt.Sprintf("Page No. %v  Options", pageNum+1),
			Items: []string{"Previous", "Next", "Exit"},
		}
	}

	_, result, err := prompt.Run()
	handle(err)

	switch result {
	case "Previous":
		// if less than 0, return to first page
		if pageNum == 0 {
			root(allTrucks, pageNum)
		} else { // else, go back one page
			root(allTrucks, pageNum-1)
		}

	case "Next":
		// if no more next pages, go to last page
		if len(allTrucks)/RESULTSPERPAGE < pageNum {
			root(allTrucks, pageNum)
		} else { // else, go to next page
			root(allTrucks, pageNum+1)
		}

	case "Exit":
		logger.Log("\nOk, goodbye! :)\n\n")
		os.Exit(0)
	}
}

func showTrucks(allTrucks []truck, pageNum int) {

	currTrucks := paginate(allTrucks, pageNum*RESULTSPERPAGE, RESULTSPERPAGE)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Truck Name",
		"Address",
		"Times",
	})

	for _, tr := range currTrucks {
		truckData := []string{
			tr.Applicant,
			tr.Location,
			fmt.Sprintf("%v  %v - %v", tr.Dayofweekstr, tr.Starttime, tr.Endtime),
		}
		table.Append(truckData)
	}

	logger.Always("There are a total of %v open food trucks in San Francisco now.", len(allTrucks))
	totalPages := len(allTrucks) / RESULTSPERPAGE
	if len(allTrucks)%RESULTSPERPAGE > 0 {
		totalPages++
	}
	logger.Always("    page %v of %v: \n\n", pageNum+1, totalPages)

	table.Render()

}
