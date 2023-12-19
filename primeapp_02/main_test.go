package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

const errMsgNotExpected = "%s got ERROR:\nexpected :\t%s \nbut got  :\t%s"
const errValueNotExpected = "%s got ERROR:\nexpected :\t%v \nbut got  :\t%v"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"seven should be prime", 7, true, fmt.Sprintf(msgIsPrime, 7)},
		{"eight shouldn't be prime", 8, false, fmt.Sprintf(msgIsNotPrime, 8, 2)},
		{"twenty-five shouldn't be prime", 25, false, fmt.Sprintf(msgIsNotPrime, 25, 5)},
		{"zero shouldn't be prime by def", 0, false, fmt.Sprintf(msgIsNotPrimeByDef, 0)},
		{"zero shouldn't be prime by def", 1, false, fmt.Sprintf(msgIsNotPrimeByDef, 1)},
		{"negative number shouldn't be prime", -11, false, fmt.Sprintf(msgNegativeNumAreNotPrimeByDef, -11)},
	}

	for _, tt := range primeTests {
		t.Run(tt.name, func(t *testing.T) {
			got, msg := isPrime(tt.testNum)
			if got != tt.expected {
				t.Errorf(errValueNotExpected, tt.name, tt.expected, got)
			}

			if tt.msg != msg {
				t.Errorf(errMsgNotExpected, tt.name, tt.msg, msg)
			}
		})
	}
}

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"seven should be prime", "7", fmt.Sprintf(msgIsPrime, 7)},
		{"eight shouldn't be prime", "8", fmt.Sprintf(msgIsNotPrime, 8, 2)},
		{"twenty-five shouldn't be prime", "25", fmt.Sprintf(msgIsNotPrime, 25, 5)},
		{"zero shouldn't be prime by def", "0", fmt.Sprintf(msgIsNotPrimeByDef, 0)},
		{"zero shouldn't be prime by def", "1", fmt.Sprintf(msgIsNotPrimeByDef, 1)},
		{"negative number shouldn't be prime", "-11", fmt.Sprintf(msgNegativeNumAreNotPrimeByDef, -11)},
		{"empty", "", msgEnterNumber},
		{"typed", "three", msgEnterNumber},
		{"decimal", "1.1", msgEnterNumber},
		{"quit", "q", ""},
		{"QUIT", "Q", ""},
		{"greek", "επτά", msgEnterNumber},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			reader := bufio.NewScanner(input)
			res, _ := checkNumbers(reader)

			if !strings.EqualFold(res, tt.expected) {
				t.Errorf("%s: expected %s, but got %s", tt.name, tt.expected, res)
			}
		})
	}
}
