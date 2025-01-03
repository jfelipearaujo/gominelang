TAG := $(shell git describe --tags --abbrev=0 2>/dev/null)
VERSION := $(shell echo $(TAG) | sed 's/v//')

install-go:
	@desiredVersion="go1.23.4"; \
	currentVersion=$$(go version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+'); \
	echo "Current version: $$currentVersion"; \
	if [ "$$desiredVersion" != "$$currentVersion" ]; then \
		echo "Installing $$desiredVersion..."; \
		sudo rm -rf /usr/local/go; \
		wget --quiet https://go.dev/dl/$$desiredVersion.linux-amd64.tar.gz; \
		sudo tar -C /usr/local -xzf $$desiredVersion.linux-amd64.tar.gz; \
		rm $$desiredVersion.linux-amd64.tar.gz; \
		export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH; \
		echo "Done!"; \
		go version; \
	else \
		echo "Go $$desiredVersion is already installed. Skipping..."; \
	fi

tag:
	@if [ -z "$(TAG)" ]; then \
        echo "No previous version found. Creating v1.0.0 tag..."; \
        git tag v1.0.0; \
    else \
        echo "Previous version found: $(VERSION)"; \
        read -p "Bump major version (M/m), minor version (R/r), or patch version (P/p)? " choice; \
        if [ "$$choice" = "M" ] || [ "$$choice" = "m" ]; then \
            echo "Bumping major version..."; \
			major=$$(echo $(VERSION) | cut -d'.' -f1); \
            major=$$(expr $$major + 1); \
            new_version=$$major.0.0; \
		elif [ "$$choice" = "R" ] || [ "$$choice" = "r" ]; then \
            echo "Bumping minor version..."; \
			minor=$$(echo $(VERSION) | cut -d'.' -f2); \
            minor=$$(expr $$minor + 1); \
            new_version=$$(echo $(VERSION) | cut -d'.' -f1).$$minor.0; \
		elif [ "$$choice" = "P" ] || [ "$$choice" = "p" ]; then \
            echo "Bumping patch version..."; \
			patch=$$(echo $(VERSION) | cut -d'.' -f3); \
            patch=$$(expr $$patch + 1); \
            new_version=$$(echo $(VERSION) | cut -d'.' -f1).$$(echo $(VERSION) | cut -d'.' -f2).$$patch; \
        else \
            echo "Invalid choice. Aborting."; \
            exit 1; \
        fi; \
        echo "Creating tag for version v$$new_version..."; \
        git tag v$$new_version; \
    fi

run:
	go run ./cmd/main.go

build:
	go build -race -o ./bin/gominelang -ldflags="-s -w -X main.GoMineLangVersion=v1.0.0" ./cmd/main.go

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -race -o ./bin/gominelang.exe -ldflags="-s -w -X main.GoMineLangVersion=v1.1.0" ./cmd/main.go

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

lint:
	golangci-lint run
