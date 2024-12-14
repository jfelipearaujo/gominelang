run:
	go run ./cmd/main.go

build:
	go build -race -o ./bin/gominelang -ldflags="-s -w" ./cmd/main.go

install: build
	go install ./cmd/main.go

test:
	go test ./internal/...

gen-mocks: ## Gen mock files using mockery
	@if command -v mockery > /dev/null; then \
		echo "Generating..."; \
		mockery; \
	else \
		read -p "Go 'mockery' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/vektra/mockery/v2@latest; \
			echo "Generating..."; \
			mockery; \
		else \
			echo "You chose not to intall mockery. Exiting..."; \
			exit 1; \
		fi; \
	fi