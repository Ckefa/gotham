all: buildcss build run runapp

buildcss:
	@npm run build:css

build:
	@go build -o bin/app cmd/main/main.go

run: 
	./bin/app

runapp:
	@go run cmd/main/main.go
