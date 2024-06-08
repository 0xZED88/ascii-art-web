package ascii_test

import (
	"testing"
	"ascii/ascii"
)

var testcases = []struct {
	name     string
	expected bool
	myString string
}{
	{"CheckNonePrintable", false, "Mohamedéé"},
	{"CheckNonePrintable", true, "Mohamed"},
}

func TestCheckNonePrintable(t *testing.T) {
	for _, char := range testcases {
		t.Run(char.name,func(t *testing.T){
			got := ascii.CheckNonePrintable(char.myString)
			if got != char.expected {
				t.Errorf("Error, Expected %t got %t", char.expected, got)
			}
		})
	}
}
	