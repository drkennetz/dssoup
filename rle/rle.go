package rle

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode is a simple string compression algorithm that compacts repeats by replacing them with their integer value
// It is a lossless compression and can be completely decoded.
func RunLengthEncode(input string) string {
	output := []string{""}
	for i := 0; i < len(input); i++ {
		// this is the current count of the char, which should start at 1
		currentCount := 1
		char := string(input[i])
		// This is called a lookahead. I'm checking the current char with the next char
		// if they are equal, I'm incrementing the current count, and the current pos
		for (i < len(input)-1) && (input[i] == input[i+1]) {
			currentCount++
			i++
		}
		// if we have more than one consecutive char, we include this in string
		// else we just include the char
		if currentCount > 1 {
			strCount := strconv.Itoa(currentCount)
			output = append(output, strCount+char)
		} else {
			output = append(output, char)
		}
	}
	return strings.Join(output, "")
}

func RunLengthDecode(input string) string {
	output := []string{""}
	var temp string
	for _, char := range input {
		if unicode.IsNumber(char) {
			temp += string(char)
		} else if (unicode.IsLetter(char) || unicode.IsSpace(char)) && len(temp) != 0 {
			times, _ := strconv.Atoi(temp)
			output = append(output, strings.Repeat(string(char), times))
			temp = ""
		} else {
			output = append(output, string(char))
		}
	}
	return strings.Join(output, "")
}
