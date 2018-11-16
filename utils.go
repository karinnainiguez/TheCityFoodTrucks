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
