BINARY_NAME=adelApp

build:
	@go mod vendor
	@echo "Building Adel..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Adel built!"

run: build
	@echo "Starting Adel..."
	@./tmp/${BINARY_NAME} &
	@echo "Adel started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Adel..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Adel!"

restart: stop start