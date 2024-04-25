package main

import (
	"fmt"
	"io"
	"net/mail"
	"os"
	"regexp"
	"sort"
)

// https://pkg.go.dev/net/mail
func main() {
	fmt.Println("Go Standard Library: net/mail")
	fmt.Println()

	filename := "testdata/mail-1.eml"
	args := os.Args[1:]
	if len(args) >= 1 {
		filename = args[0]
	}

	file, err := os.Open(filename)
	if err != nil {
		handleError(fmt.Sprintf("Error opening email file: %s", filename), err)
	}
	defer file.Close()

	message, err := mail.ReadMessage(file)
	if err != nil {
		handleError("Error reading message", err)
	}
	printMessage(message)

	fmt.Println("Creating blacklist for headers to be removed...")
	blacklist := []string{"X-Ms", "^Thread-.*", ".*-Language$"}

	fmt.Println("Creating custom headers to be added...")
	custom := make(mail.Header)
	custom["X-Acronis-1"] = []string{"First custom header"}
	custom["X-Acronis-2"] = []string{"Second custom header...", "...with two values"}
	custom["X-Acronis-3"] = []string{"Third custom header...", "...with three values...", "...third value."}

	fmt.Println("Modifying headers...")
	err = modifyHeaders(message, blacklist, custom)
	if err != nil {
		handleError("Error stripping and appending headers", err)
	}
	printHeader(message)

	fmt.Println("Done.")
}

// Strips out headers that match any of the regex patterns in the blacklist and appends custom headers
func modifyHeaders(message *mail.Message, blacklist []string, custom mail.Header) error {
	header := make(mail.Header)
	for key, value := range message.Header {
		matches := false
		for _, pattern := range blacklist {
			m, err := regexp.MatchString(pattern, key)
			if err != nil {
				return err
			}
			if m {
				//fmt.Printf("Key %s matches pattern: %s\n", key, pattern)
				matches = true
			}
		}
		if !matches {
			header[key] = value
		}
	}
	for key, value := range custom {
		header[key] = value
	}
	message.Header = header
	return nil
}

func printMessage(message *mail.Message) {
	fmt.Println("============ Message ============")
	printHeader(message)
	printBody(message)
	fmt.Println("=================================")
}

func printHeader(message *mail.Message) {
	fmt.Println("------------ Header ------------")
	fmt.Println("Common headers:")
	from, err := mail.ParseAddress(message.Header.Get("From"))
	if err != nil {
		handleError("Error reading 'From' header", err)
	}
	fmt.Printf("  From: %v\n", from)

	toList, err := mail.ParseAddressList(message.Header.Get("To"))
	if err != nil {
		handleError("Error reading 'To' header", err)
	}
	for _, addr := range toList {
		fmt.Printf("  To: %v\n", addr.String())
	}

	subject := message.Header.Get("Subject")
	fmt.Printf("  Subject: %v\n", subject)

	fmt.Println("All headers:")
	keys := make([]string, 0, len(message.Header))
	for key := range message.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("  %s:%s\n", key, message.Header[key])
	}
	fmt.Println()
}

func printBody(message *mail.Message) {
	fmt.Println("------------ Body ------------")
	body, err := io.ReadAll(message.Body)
	if err != nil {
		handleError("Error reading body", err)
	}
	fmt.Printf("%v\n", string(body))
	fmt.Println()
}

func handleError(message string, err error) {
	fmt.Print(message)
	fmt.Printf(" Error: %v\n", err)
	os.Exit(1)
}
