package main

import (
	"fmt"
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
