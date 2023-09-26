.PHONY: vendor run traffic

vendor:
	go mod tidy
	go mod vendor

run:
	go run cmd/zipdb/*.go

traffic:
	echo "GET http://localhost:8080/90210" | vegeta attack -duration=5s | tee results.bin | vegeta report

.PHONY: seed
seed:
	go run cmd/seed/*.go

create:
	@curl -i -d @data/12345.json \
		-H "Content-Type: application/json" \
		-X POST http://localhost:8080/

ZIP ?= 90210
read:
	@curl -s http://localhost:8080/$(ZIP)

update:
	@curl -i -d @data/90210.json \
		-H "Content-Type: application/json" \
		-X PUT http://localhost:8080/90210

delete:
	@curl -i -X DELETE http://localhost:8080/12345