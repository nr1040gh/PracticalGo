package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

// inputs satisfy the Reader and Writer interfaces
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"

	// look up what's special about Fprintf
	// Fprintf writes formatted message to writer
	// Sprintf returns a string in the given format
	// Printf writes to standard output and returns number of bytes and error
	fmt.Fprintf(w, msg)

	// look up documentation for NewScanner and Scan
	scanner := bufio.NewScanner(r)
	scanner.Scan() //Scan() returns once the new line character is read by default
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text() // Text() returns the read data as a string
	if len(name) == 0 {
		return "", errors.New("you didn't enter your name")
	}

	return name, nil
}

type config struct {
	numTimes int
	// printUsage bool
}

// Command-line arguments supplied to a program are available via the Args slice defined in the os package
func parseArgs(w io.Writer, args []string) (config, error) {
	// var numTimes int
	// var err error
	c := config{}

	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() != 0 {
		return c, errors.New("positional arguments specified")
	}
	return c, nil

	// if len(args) != 1 {
	// 	// error strings should not be capitalized
	// 	return c, errors.New("invalid number of arguments")
	// }

	// // default value for bool is false?
	// if args[0] == "-h" || args[0] == "--help" {
	// 	c.printUsage = true
	// 	return c, nil
	// }

	// // strconv.Atoi converts str to int
	// numTimes, err = strconv.Atoi(args[0])
	// if err != nil {
	// 	return c, err
	// }
	// c.numTimes = numTimes

	// return c, nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}

	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
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
	c, err := parseArgs(os.Stderr, os.Args[1:])
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
