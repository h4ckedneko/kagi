.PHONY: test bench format

test:
	@ go test -v -race

bench:
	@ go test -bench=. -benchmem

format:
	@ go fmt ./...
