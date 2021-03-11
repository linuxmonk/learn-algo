.PHONY: all

all: build test

build:
	go build .

test:
	go test -v .

bench:
	V=1 go test -run=zzz -bench=BubbleSort -count=5 | tee old.txt
	V=2 go test -run=zzz -bench=BubbleSort -count=5 | tee new.txt
	@echo ">> ======================================"
	@echo ">> Comparing statistics between v1 and v2"
	@echo ">> ======================================"
	benchstat old.txt new.txt

cpuprofile:
	go test -run=zzz -bench=BubbleSort -count=5 -cpuprofile cpuprof.out

clean:
	rm -f old.txt new.txt sorting.test cpuprof.out
