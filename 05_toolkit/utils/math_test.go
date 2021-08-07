package utils

import "testing"

func TestAdd(t *testing.T) {
	actual := Sum(12, 4, 4)
	expected := 20

	if actual != expected {
		t.Errorf("Invalid value. Expected: %d, got: %d", expected, actual)
	}
}

func TestSubtract(t *testing.T) {
	actual := Subtract(12, 2)
	expected := 10

	if actual != expected {
		t.Errorf("Invalid value. Expected: %d, got: %d", expected, actual)
	}
}

func TestDivision(t *testing.T) {
	actual := Divide(12, 2)
	expected := 6.0

	if actual != expected {
		t.Errorf("Invalid value. Expected: %f, got: %f", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	if Multiply(12, 2) != 24 {
		t.Errorf("Invalid value. Expected: %d, got: %d", 24, Multiply(12, 2))
	}
}
