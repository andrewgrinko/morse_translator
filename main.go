package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Decoded is the result of translation
type Decoded struct {
	Data string
}

func main() {
	http.HandleFunc("/", getRequest)
	http.HandleFunc("/translate", postRequest)

	log.Println("Listening on http://localhost:1337/")
	log.Fatal(http.ListenAndServe(":1337", nil))
}

// ToMorse translates to morse
func ToMorse(s string) (result string) {
	letters := strings.Split(s, "")

	for _, v := range letters {
		if code, ok := MorseValues[v]; ok {
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
			for key, code := range MorseValues {
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

func sendTemplate(w http.ResponseWriter, r *http.Request, d *Decoded) {
	t := template.New("index")
	t, err := t.ParseFiles("index.html")
	if err != nil {
		fmt.Printf("Error while parsing template: %v", err)
	}
	t.ExecuteTemplate(w, "index.html", d)
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	d := &Decoded{Data: ""}
	sendTemplate(w, r, d)
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	var result Decoded
	data := r.PostFormValue("data")
	data = strings.ToLower(data)

	// create regexp to check if its morse or english
	pattern := "[a-zA-Z]"
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	if re.MatchString(data) {
		result.Data = ToMorse(data)
	} else {
		result.Data = strings.ToUpper(FromMorse(data))
	}

	sendTemplate(w, r, &result)
}
