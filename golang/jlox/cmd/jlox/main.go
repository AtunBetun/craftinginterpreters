package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"atunbetun.jlox.com/pkg/logs"
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
	logger.Debug(fmt.Sprintf("opening file: %s", path))
	file, err := os.Open(path)

	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	defer file.Close()
	b, err := io.ReadAll(file)
	logger.Debug(string(b))
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
		runFile(os.Args[1], logger)
	} else {
		logger.Debug("prompt mode")
	}
}
