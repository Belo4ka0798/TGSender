###########################################
# MakeFile for "Server TGSEnder"
###########################################

.PHONY: all
all: build

.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: build
build:
	go build -o tgsender cmd/app/main.go

.PHONY: rebuild
rebuild: clean build

.PHONY: clean
clean:
	go clean
	rm tgsender

.SILENT: build clean