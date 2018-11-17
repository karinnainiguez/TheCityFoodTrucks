package main

import (
	"os"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

func handle(err error) {
	if err != nil {
		logger.Critical("%s\n", err.Error())
		os.Exit(1)
	}

}

func paginate(trucks []truck, skip int, size int) []truck {
	if skip > len(trucks) {
		skip = len(trucks)
	}

	end := skip + size
	if end > len(trucks) {
		end = len(trucks)
	}

	currentCollection := trucks[skip:end]
	return currentCollection
}
