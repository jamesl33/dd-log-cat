package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
)

// exit if the given error is non-nil.
func exit(err error) {
	if err == nil {
		return
	}

	defer os.Exit(1)

	fmt.Printf("Error: %s\n", err)
}

// main accepts stdin (or a file path) and extracts the JSON message body from it.
func main() {
	switch len(os.Args) {
	case 1:
		err := dump(os.Stdin)
		exit(err)
	case 2:
		file, err := os.Open(os.Args[1])
		exit(err)
		defer file.Close()

		err = dump(file)
		exit(err)
	default:
		exit(fmt.Errorf("%s [string]", os.Args[0]))
	}
}

// dump the 'Message' attribute from the given DataDog CSV export.
func dump(r io.Reader) error {
	reader := csv.NewReader(r)

	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	const msg = "Message"

	// Find the column index for the JSON message
	idx := slices.Index(header, msg)

	if idx < 0 {
		return fmt.Errorf("input doesn't have a %q column", msg)
	}

	for {
		line, err := reader.Read()

		// Read to the end of the file, exit
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return fmt.Errorf("failed to read line: %w", err)
		}

		// Ignore log lines which aren't JSON formatted
		if !json.Valid([]byte(line[idx])) {
			continue
		}

		fmt.Println(line[idx])
	}

	return nil
}
