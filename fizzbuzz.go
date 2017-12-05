package fizzbuzz

import "strconv"

func fizzbuzz(x int) string {
	if x%15 == 0 {
		return "FizzBuzz"
	} else if x%5 == 0 {
		return "Buzz"
	} else if x%3 == 0 {
		return "Fizz"
	} else {
		return strconv.Itoa(x)
	}
}
