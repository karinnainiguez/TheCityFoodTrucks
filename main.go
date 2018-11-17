package main

import (
	"github.com/kubicorn/kubicorn/pkg/logger"
)

func main() {
	// obtain collection of available food trucks
	truckCollection := getQuery()

	// If results, show.
	if len(truckCollection) > 0 {
		root(truckCollection, 0)
	} else {
		logger.Critical("\nSo sorry!  It looks like there are no available food trucks right now in the city :(\n")
	}

}
