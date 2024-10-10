package main

import (
	"bytes"
	"errors"
	"testing"
)

// config doesnt have a type associated with it?
type testConfig struct {
	args []string
	err  error
	c    config
}

// testing module identifies functions that start with Test?
func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		// {
		// 	args: []string{"-h"},
		// 	err:  nil,
		// 	c:    config{printUsage: true, numTimes: 0}, // composite literals require , at the end?
		// },
		{
			args: []string{"-h"},
			err:  errors.New("flag: help requested"),
			c:    config{numTimes: 0},
		},
		{
			args: []string{"-n", "10"},
			err:  nil,
			c:    config{numTimes: 10},
		},
		{
			args: []string{"-n", "abc"},
			err:  errors.New("invalid value \"abc\" for flag -n: parse error"),
			c:    config{numTimes: 0},
		},
		{
			args: []string{"-n", "1", "foo"},
			err:  errors.New("positional arguments specified"),
			c:    config{numTimes: 1},
		}, // composite literals require , at the end?
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		c, err := parseArgs(byteBuf, tc.args)
		// why do we use err.Error()? check documentation
		// uses Error() interface! https://stackoverflow.com/a/42770341
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		// if c.printUsage != tc.c.printUsage {
		// 	t.Errorf("Expected printUsage to be: %v, got: %v\n", tc.c.printUsage, c.printUsage)
		// }
		if c.numTimes != tc.c.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.c.numTimes, c.numTimes)
		}
		byteBuf.Reset()
	}
}
