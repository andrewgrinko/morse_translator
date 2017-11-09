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
		fmt.Println(fromMorse(message))
	}

	if isToMorse {
		fmt.Println(toMorse(message))
	}
}

func toMorse(s string) (result string) {
	letters := strings.Split(s, "")

	for _, v := range letters {
		if code, ok := morseValues[v]; ok {
			result += code
			result += " "
		} else {
			result += "| "
		}
	}

	return
}

func fromMorse(s string) (result string) {
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

	return
}
