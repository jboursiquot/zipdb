.PHONY: vendor run traffic

vendor:
	go mod tidy
	go mod vendor

run:
	go run cmd/zipdb/*.go

traffic:
	echo "GET http://localhost:8080/90210" | vegeta attack -duration=5s | tee results.bin | vegeta report