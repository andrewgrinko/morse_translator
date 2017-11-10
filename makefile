dev:
	go run main.go morse_dict.go

build:
	go build -o ./morse_translator

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./morse_translator
