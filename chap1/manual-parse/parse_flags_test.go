package main

import (
	"errors"
	"testing"
)

// config doesnt have a type associated with it?
type testConfig struct {
	args []string
	err  error
	config
}

// testing module identifies functions that start with Test?
func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUsage: true, numTimes: 0}, // composite literals require , at the end?
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUsage: false, numTimes: 10},
		},
		{
			args:   []string{"abc"},
			err:    errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			config: config{printUsage: false, numTimes: 0},
		},
		{
			args:   []string{"1", "foo"},
			err:    errors.New("Invalid number of arguments"),
			config: config{printUsage: false, numTimes: 0},
		}, // composite literals require , at the end?
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		// why do we use err.Error()? check documentation
		// uses Error() interface! https://stackoverflow.com/a/42770341
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if c.printUsage != tc.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n", tc.printUsage, c.printUsage)
		}
		if c.numTimes != tc.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.numTimes, c.numTimes)
		}
	}
}
