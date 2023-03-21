APP_NAME=card-api
BUILD_DIR=build
SRC_DIR=.

.PHONY: clean build run
build:	clean
	go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)/main.go handler.go deck.go card.go
run: build
	./$(BUILD_DIR)/$(APP_NAME)
test: 
	go test
