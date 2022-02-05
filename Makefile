build:
	go build -o bin/webchat.exe cmd/app/main.go

run: build
	./bin/webchat.exe
