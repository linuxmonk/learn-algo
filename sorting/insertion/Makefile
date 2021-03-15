.PHONY: all

all: build test

build:
	go build .

test:
	go test -v .

bench:
	go test -run=zzz -bench=InsertionSort -count=5

cpuprofile:
	go test -run=zzz -bench=InsertionSort -count=5 -cpuprofile cpuprof.out

clean:
	rm -f insertion.test puprof.out
