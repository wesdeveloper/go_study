package main

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(2, 4)
	expected := 6

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestSubtract(t *testing.T) {
	sum := Subtract(2, 4)
	expected := -2

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestMultiply(t *testing.T) {
	sum := Multiply(2, 4)
	expected := 8

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestDivide(t *testing.T) {
	sum := Divide(4, 2)
	expected := 2

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}