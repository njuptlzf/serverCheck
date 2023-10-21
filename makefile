.PHONY: gen_checker
gen_checker:
	bash scripts/gen_code.sh ${name}

.PHONY: build
build:
	go build -o _output/bin/servercheck cmd/main.go
	chmod +x _output/bin/servercheck

.PHONY: test
test:
	go test ./...
