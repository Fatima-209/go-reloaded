package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Base num: 10 = decimal; 2 = binary; 8 = octal; 16 = hexaDecimal
func hexToDec(hexStr string) string {
	// bit size of the resulting int 64
	dec, er := strconv.ParseInt(hexStr, 16, 64)
	if er != nil {
		fmt.Println("error: invalid hexadecimal string", er)
		return ""
	}
	// FormatInt (int (dec)> string (10 is the base for the output string
	return strconv.FormatInt(dec, 10)
}
func binToDec(binStr string) string {
	dec, er := strconv.ParseInt(binStr, 2, 64)
	if er != nil {
		fmt.Println("error: invalid binary string", er)
		return ""
	}
	return strconv.FormatInt(dec, 10)
}
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
}
func findCountLow(word string) int {
	start := strings.Index(word, "(low,") + 1
	end := strings.Index(word, ")")
	count, _ := strconv.Atoi(word[start:end])
	return count
}
func findCountUp(word string) int {
	start := strings.Index(word, "(up,") + 1
	end := strings.Index(word, ")")
	count, _ := strconv.Atoi(word[start:end])
	return count
}
func findCountCap(word string) int {
	start := strings.Index(word, "(cap,") + 1
	end := strings.Index(word, ")")
	count, _ := strconv.Atoi(word[start:end])
	return count
}
