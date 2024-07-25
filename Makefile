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
	/bin/rm -f cgmsort cgmsort.test
	/bin/rm -f cpu.prof mem.prof profile001.callgraph.out

profile:
	go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

test:
	go test

