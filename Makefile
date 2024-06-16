.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt main.go
	go fmt mergesort.go
	go fmt randomdata.go

vet: fmt
	go vet main.go mergesort.go

build: vet
	go build -o cgmsort main.go mergesort.go

clean:
	/bin/rm -f cgmsort

test:
	go test

