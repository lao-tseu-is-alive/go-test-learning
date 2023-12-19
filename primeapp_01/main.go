package main

import "fmt"

const (
	msgIsNotPrimeByDef             = "%d is not prime, by definition."
	msgNegativeNumAreNotPrimeByDef = "%d is negative number so it's not prime, by definition."
	msgIsPrime                     = "%d is a prime number."
	msgIsNotPrime                  = "%d is not a prime number because it is divisible by %d."
)

func main() {
	n := 2

	_, msg := isPrime(n)
	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf(msgIsNotPrimeByDef, n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, fmt.Sprintf(msgNegativeNumAreNotPrimeByDef, n)
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf(msgIsNotPrime, n, i)
		}
	}

	return true, fmt.Sprintf(msgIsPrime, n)
}
