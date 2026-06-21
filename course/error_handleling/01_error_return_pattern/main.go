package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// go dont use exceptions for normal failures
	// functions -> return errors as normal return values
	// this is called error return pattern

	// val, err := someFunc()
	// if err != nil {handle error}

	res := run()

	fmt.Println(res)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := "3"

	level, err := parseLevel(input)

	if err != nil {
		return err
	}

	fmt.Println("Selected level:", level)

	return nil
}

func parseLevel(s string) (int, error) {
	// Error return pattern
	// nil error means no error
	// non nil error means there is an error

	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("Level must be a number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("Level must be 1 to 5")
	}

	return n, nil
}
