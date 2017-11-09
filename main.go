package main

import (
	"fmt"
	"strings"
)

var morseValues = map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"h": "....",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"p": ".--.",
	"q": "--.-",
	"r": ".-.",
	"s": "...",
	"t": "-",
	"u": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",
}

func main() {
	message := "... --- ... | .... . .-.. .--."
	isToMorse := false
	isFromMorse := true

	if isFromMorse {
		fmt.Println(FromMorse(message))
	}

	if isToMorse {
		fmt.Println(ToMorse(message))
	}
}

// ToMorse translates to morse
func ToMorse(s string) (result string) {
	letters := strings.Split(s, "")

	for _, v := range letters {
		if code, ok := morseValues[v]; ok {
			result += code
			result += " "
		} else {
			result += "| "
		}
	}

	result = strings.TrimSpace(result)

	return
}

// FromMorse translates from morse
func FromMorse(s string) (result string) {
	words := strings.Split(s, "|")

	for _, word := range words {
		var decodedWord string
		letters := strings.Split(word, " ")

		for _, letter := range letters {
			for key, code := range morseValues {
				if code == letter {
					decodedWord += key
				}
			}
		}

		result += decodedWord
		result += " "
	}

	result = strings.TrimSpace(result)

	return
}
