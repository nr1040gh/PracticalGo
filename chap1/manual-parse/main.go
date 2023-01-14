package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// inputs satisfy the Reader and Writer interfaces
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"

	// look up what's special about Fprintf
	fmt.Fprintf(w, msg)

	// look up documentation for NewScanner and Scan
	scanner := bufio.NewScanner(r)
	scanner.Scan() //Scan() returns once the new line character is read by default
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text() // Text() returns the read data as a string
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}

	return name, nil
}

type config struct {
	numTimes   int
	printUsage bool
}

// Command-line arguments supplied to a program are available via the Args slice defined in the os package
func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}

	// default value for bool is false?
	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}

	// strconv.Atoi converts str to int
	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes

	return c, nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) && !(c.printUsage) {
		return errors.New("Must specify a number greater than 0")
	}

	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)

	return nil
}

func greetUser(c config, name string, w io.Writer) {
	// Check the difference between Sprintf and Fprintf
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

// Whats the difference between using “ and ” or ""
var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
A greeter application which prints the name your entered <integer> number of times.
`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

// build with: go build -o application
func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
