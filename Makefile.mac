BINARY_NAME=adelApp

build:
	@go mod vendor
	@echo "Building adel..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "adel built!"

run: build
	@echo "Starting adel..."
	@./tmp/${BINARY_NAME} &
	@echo "adel started!"

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
	@echo "Stopping adel..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped adel!"

restart: stop start