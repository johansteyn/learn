package main

//import log "github.com/sirupsen/logrus"
import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Go Standard Library: log")
	fmt.Println()

	// Create a default logger that outputs to stdout
	logger := log.Default()
	logger.Println("Default logger's Println function.")
	logger.Print("Default logger's Print function.")
	logger.Printf("Default logger's Printf %s.", "function")
	// Q: So what is the difference between Print and Println?
	// A: https://groups.google.com/g/golang-nuts/c/SOibvyF6Tjo

	// The Fatal function is like Println + os.Exit(1)
	//logger.Fatal("Default logger's Fatal function.")

	// The Panic function is like Println + panic
	//logger.Panic("Default logger's Panic function.")

	// Can set a prefix
	logger.SetPrefix("prefix: ")
	logger.Println("Default logger's Println function with prefix")

	// Can set flags
	logger.SetFlags(log.Lmicroseconds | log.Lmsgprefix)
	logger.Println("Default logger's Println function with flags set")

	// Can log to a file instead of stdout
	file, err := os.Create("log.txt")
	if err != nil {
		logger.Fatal("Error creating log file")
	}
	logger.SetOutput(file)
	logger.Println("Default file logger's Println function.")

	// Or create a custome logger that writes to a file and includes a prefix
	logger = log.New(file, "prefix: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	logger.Println("Custom file + prefix logger's Println function.")

	fmt.Println()
}
