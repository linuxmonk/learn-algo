.PHONY: all

all: build test

build:
	go build .

test:
	go test -v .

bench:
	go test -run=zzz -bench=Largest -count=3

cpuprofile:
	go test -run=zzz -bench=Largest -cpuprofile cpuprof.out

memprofile:
	go test -run=zzz -bench=Largest -benchmem -memprofile m.out

clean:
	rm -f largest.test cpuprof.out m.out
