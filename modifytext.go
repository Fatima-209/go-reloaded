package main

import "strings"

func modifyText(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		switch {
		case i > 0 && (strings.HasSuffix(words[i], "(hex)") || strings.HasSuffix(words[i], "(HEX)") || strings.HasSuffix(words[i], "(Hex)")):
			words[i-1] = hexToDec(words[i-1])
			// the ... are used to add each item individually, making it in the right order
			words = append(words[:i], words[i+1:]...)
			// The loop decremented to accout for the the removed element (hex)
			i--
		case i > 0 && (strings.HasSuffix(words[i], "(bin)") || strings.HasSuffix(words[i], "(BIN)") || strings.HasSuffix(words[i], "(Bin)")):
			words[i-1] = binToDec(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case i > 0 && (strings.HasSuffix(words[i], "(up)") || strings.HasSuffix(words[i], "(UP)") || strings.HasSuffix(words[i], "(Up)")):
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case i > 0 && (strings.HasSuffix(words[i], "(low)") || strings.HasSuffix(words[i], "(LOW)") || strings.HasSuffix(words[i], "(Low)")):
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case i > 0 && (strings.HasSuffix(words[i], "(cap)") || strings.HasSuffix(words[i], "(CAP)") || strings.HasSuffix(words[i], "(Cap)")):
			words[i-1] = capitalize(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		case i >= 0 && (strings.HasPrefix(words[i], "(low,") || strings.HasPrefix(words[i], "(LOW,") || strings.HasPrefix(words[i], "(Low,")):
			count := findCountLow(words[i+1])
			if i-count >= 0 {
				for x := i - count; x < i; x++ {
					words[x] = strings.ToLower(words[x])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		case i >= 0 && (strings.HasPrefix(words[i], "(up,") || strings.HasPrefix(words[i], "(UP,") || strings.HasPrefix(words[i], "(Up,")):
			count := findCountUp(words[i+1])
			if i-count >= 0 {
				for x := i - count; x < i; x++ {
					words[x] = strings.ToUpper(words[x])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		case i > 0 && (strings.HasPrefix(words[i], "(cap,") || strings.HasPrefix(words[i], "(CAP,") || strings.HasPrefix(words[i], "(Cap,")):
			count := findCountCap(words[i+1])
			if i-count >= 0 {
				for x := i - count; x < i; x++ {
					words[x] = capitalize(words[x])
				}
			}
			words = append(words[:i], words[i+2:]...)
			i--
		case words[i] == "." || words[i] == "," || words[i] == "!" || words[i] == "?" ||
			words[i] == ":" || words[i] == ";":
			words[i-1] = words[i-1] + words[i]
			words = append(words[:i], words[i+1:]...)
			i--
		case i > 0 && (words[i] == "..." || words[i] == "!?"):
			words[i-1] = words[i-1] + words[i]
			words = append(words[:i], words[i+1:]...)
			i--
		case words[i] == "'":
			j := i + 1
			for j < len(words) && words[j] != "'" {
				j++
			}
			if j < len(words) {
				words[i+1] = words[i] + words[i+1]
				words[j-1] = words[j-1] + words[j]
				// Gives us the slice from the beggining until the word before the first single quote
				// And the word after the first single quote till the word before the sec single quote
				words = append(words[:i], words[i+1:j]...)
				// From before the 2nd single quote untill the end
				words = append(words[:j-1], words[j+1:]...)
				j -= 2
			}
			i = j
		// [0] for the first letter of the next word
		case (words[i] == "a" || words[i] == "A") && i < len(words)-1 && strings.ContainsAny(string(words[i+1][0]), "aeiouhAEIOUH"):
			words[i] += "n"
		}
	}
	return strings.Join(words, " ")
}
