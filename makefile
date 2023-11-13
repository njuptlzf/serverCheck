.PHONY: gen_checker
gen_checker:
	bash scripts/gen_code.sh ${name}

.PHONY: build
build:
	go build -o _output/bin/servercheck cmd/main.go

.PHONY: test
test:
	go test ./...
