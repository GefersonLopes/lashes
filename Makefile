APP_NAME := emilly-lashes

run:
	go run main.go

build:
	go build -o $(APP_NAME)

clean:
	rm -f $(APP_NAME)

migrate:
	go run main.go -migrate

.PHONY: run build clean migrate
