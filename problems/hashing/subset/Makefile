.PHONY: all

all: build test

build:
	go build .

test:
	go test -v .

bench:
	V=1 go test -run=zzz -bench=Subset -count=5 | tee old.txt
	V=2 go test -run=zzz -bench=Subset -count=5 | tee new.txt
	@echo ">> ======================================"
	@echo ">> Comparing statistics between v1 and v2"
	@echo ">> ======================================"
	benchstat old.txt new.txt

cpuprofile:
	V=1 go test -run=zzz -bench=Subset -count=5 -cpuprofile cpuprof1.out
	V=2 go test -run=zzz -bench=Subset -count=5 -cpuprofile cpuprof2.out

clean:
	rm -f old.txt new.txt hashing.test cpuprof1.out cpuprof2.out
