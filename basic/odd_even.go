package basic

import "fmt"

// prime number functioms
func OddEven(number int) {
	// # this function will list odd or even numbers till a some number
	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println("Even number found ", i)
		} else {
			fmt.Println("Odd number found  ", i)
		}
	}

}

// prime number functioms
func ListPrimeNumbers(n int) {

	for number := 2; number <= n; number++ {
		isPrime := true
		if number != 0 {
			if number == 2 {
				fmt.Println("Found prime number ", number)
			} else {

				for j := 2; j <= (number / 2); j++ {
					if number%j == 0 {
						isPrime = false
						// break
					}

				}
				if isPrime == true {
					fmt.Println("Found prime number ", number)
				}

			}

		}
	}

}
