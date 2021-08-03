package utils

import (
	"testing"
)

func TestMakeExcited(t *testing.T) {
	if MakeExcited("omg so exciting!") != ("OMG SO EXCITING!") {
		t.Errorf("make excited function was incorrect! Expected: %s, Actual: %s", "OMG SO EXCITING!", MakeExcited("omg so exciting!"))
	}
}
