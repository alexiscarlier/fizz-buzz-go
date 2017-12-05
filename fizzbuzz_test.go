package fizzbuzz

import (
	"testing"
)

func TestFizz(t *testing.T) {
	value := fizzbuzz(3)
	if value != "Fizz" {
		t.Errorf("Value was incorrect")
		// t.Errorf("Value was incorrect, got: %s, want %s", value, "Fizz")
	}
}

func TestBuzz(t *testing.T) {
	value := fizzbuzz(5)
	if value != "Buzz" {
		t.Errorf("Value was incorrect")
		// t.Errorf("Value was incorrect, got: %s, want %s", value, "Fizz")
	}
}

func TestFizzBuzz(t *testing.T) {
	value := fizzbuzz(15)
	if value != "FizzBuzz" {
		t.Errorf("Value was incorrect")
		// t.Errorf("Value was incorrect, got: %s, want %s", value, "Fizz")
	}
}

func TestNumber(t *testing.T) {
	value := fizzbuzz(16)
	if value != "16" {
		t.Errorf("Value was incorrect")
		// t.Errorf("Value was incorrect, got: %s, want %s", value, "Fizz")
	}
}
