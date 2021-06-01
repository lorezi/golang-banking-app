APP?=banking_app


.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -o ${APP} main.go

.PHONY: run
## run: runs the go run build-binary
run: build
	./banking_app

.PHONY: watch
## watch: watch the project for go file changes
watch:
	ulimit -n 1000 
	reflex -s -r '\.go$$' make run

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning..."
	@go clean

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...


.PHONY: start-db
## start-db: to start mysql db instance
start-db:
	cd ./docker && docker-compose up --build

.PHONY: stop-db
## stop-db: to stop mysql db instance
stop-db:
	cd ./docker && docker-compose down
	

.PHONY: docker-compose-up
## docker-compose-up: to spin up multiple services
docker-compose-up:
	docker-compose up --build

.PHONY: docker-compose-down
## docker-compose-down: to stop and remove unwanted image builds
docker-compose-down:
	docker-compose down
	docker system prune --volumes --force


.PHONY: help
## help: prints this help message
help:

	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'