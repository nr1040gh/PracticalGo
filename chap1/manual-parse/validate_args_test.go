package main

import (
	"errors"
	"testing"
)

// testing module identifies functions that start with Test?
func TestValidateArgs(t *testing.T) {
	tests := []testConfig{
		{
			c:   config{}, // composite literals require , at the end?
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: -1},
			err: errors.New("Must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: 10},
			err: nil,
		},
		{
			c:   config{printUsage: true},
			err: nil,
		}, // composite literals require , at the end?
	}

	for _, tc := range tests {
		err := validateArgs(tc.c)
		// why do we use err.Error()? check documentation
		// uses Error() interface! https://stackoverflow.com/a/42770341
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
	}
}
