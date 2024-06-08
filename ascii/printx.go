package ascii

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func CheckNonePrintable(s string) bool {
	for _, rn := range s {
		if !(rn >= 32 && rn <= 126) {
			return false
		}
	}
	return true
}

func PrintX(s string, tabl2D [][]string, result *string) {
	for i := 0; i < 8; i++ {
		for j := 0; j < len(s); j++ {
			if CheckNonePrintable(string(s[j])) {
				*result += tabl2D[s[j]-32][i]
			}
		}
		*result += "<br>"
	}
}

// var time int
func Batata(fileName string, txt string) (result string, err bool) {
	// passing File Argument, and getting the data to be processed
	table, err := ReadFile(fileName)
	if err {
		return
	}

	text := strings.Split(txt, "\r\n") // "" "g"
	// Convert Our File to 2D Table
	tabl2D := AddingData(table)

	for _, char := range text {
		if char == "" {
			result += "<br>"
		} else {
			PrintX(char, tabl2D, &result)
		}
	}
	return
}
// checking if the slice Empty or not
func IsEmpty(tab []string) bool {
	for _, char := range tab {
		if len(char) != 0 {
			return false
		}
	}
	return true
}

// turning Data to a 2D Table
func AddingData(table []string) [][]string {
	tabl2D := make([][]string, len(table))
	for i := 0; i < len(table); i++ {
		tabl2D[i] = strings.Split(table[i], "\n")
	}
	return tabl2D
}

// Reading the Ascii character file got
func ReadFile(fileName string) ([]string, bool) {
	input, err := os.ReadFile(fileName + ".txt")
	if err != nil {
		fmt.Println("File Not Found")
		return []string{}, true
	}
	var tab []string
	pattern := regexp.MustCompile(`\r\n`)
	v := pattern.ReplaceAllString(string(input), "\n")
	tab = strings.Split(string(v[1:]), "\n\n")
	return tab, false
}
