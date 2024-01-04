package main

//import log "github.com/sirupsen/logrus"
import (
	"fmt"
	"log/slog"
)

func main() {
	fmt.Println("Go Standard Library: slog")
	fmt.Println()

	// Create a default logger that outputs to stdout
	logger := slog.Default()
	logger.Info("Default logger's Info function.")
	// Doesn't log anything at Debug level - need to figure out how to set the log level...
	logger.Debug("Default logger's Debug function.")
	logger.Warn("Default logger's Warn function.")
	logger.Error("Default logger's Error  function.")

	// TODO: Much more to explore...

	fmt.Println()
}
