package main

import (
	"atunbetun.jlox.com/pkg/logs"
	"fmt"
	"log/slog"
	"os"
)

func argumentPrint(logger *slog.Logger) {
	x := 0
	for _, val := range os.Args {
		logger.Debug(fmt.Sprintf("x: %d", x))
		logger.Debug(val)
	}

}

// TODO: should be on pkg once abstraction is found
func runFile(path string, logger *slog.Logger) {
	file, err := os.Open(path)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	defer file.Close()
}

func main() {
	logger := logs.CreateLogger()
	logger.Debug(fmt.Sprintf("length: %d", len(os.Args)))

	argumentPrint(logger)

	if len(os.Args) > 2 {
		logger.Info("Usage: jlox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		logger.Debug("running jlox")
		logger.Info(fmt.Sprintf("argument: %s", os.Args[1]))
	} else {
		logger.Info("doing nothing")
	}
}
