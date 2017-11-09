package main

import "testing"

var morseMessage = "... --- ..."
var englishMessage = "sos"

func TestToMorse(t *testing.T) {
	t.Log("Testing translation of SOS from English into Morse")
	if result := ToMorse(englishMessage); result != morseMessage {
		t.Errorf("Expected morse message to be %v, but it was %v", morseMessage, result)
	}
}

func TestFromMorse(t *testing.T) {
	t.Log("Testing translation of SOS from Morse into English")
	if result := FromMorse(morseMessage); result != englishMessage {
		t.Errorf("Expected english message to be %v, but it was %v", englishMessage, result)
	}
}
