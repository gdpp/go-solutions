package main

import (
	"errors"
	"fmt"
)

func main() {
	// defer resp.body.Close()

}

func doWork(succ bool) error {
	// resource related
	// start message -> resource -> aquire
	// cleanup message -> resource -> released

	fmt.Println("Starting work: resource aquire")

	defer fmt.Println("Finished work: resource released")

	if !succ {
		return errors.New("something went wrong")
	}

	fmt.Println("Doing work")
	fmt.Println("Work done")

	return nil
}
