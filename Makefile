## benchmark: 📈 Benchmark code performance
benchmark:
	go test ./... -benchmem -bench=. -run=^Benchmark_$

## test: 🚦 Execute all tests
test:
	go run gotest.tools/gotestsum@latest -f testname -- ./... -race -count=1 -shuffle=on

## tidy: 📌 Clean and tidy dependencies
tidy:
	go mod tidy -v