.PHONY: mock test


stub:
	@echo "It's stub"

mock:
	go generate -v ./...

test:
	go test -v tests/*