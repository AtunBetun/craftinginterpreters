package main

import (
	"fmt"
	"log/slog"
	"os"
)

func createLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}

func main() {
	logger := createLogger()
	logger.Debug(fmt.Sprintf("length: %d", len(os.Args)))

	x := 0
	for _, val := range os.Args {
		logger.Debug(fmt.Sprintf("x: %d", x))
		logger.Debug(val)
	}

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

func runFile(path string) {
	logger := createLogger()
	file, err := os.Open(path)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	defer file.Close()
}
