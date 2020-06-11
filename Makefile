build:
	go build -o uuid main.go

install: build
	sudo cp uuid /usr/local/bin/uuid
