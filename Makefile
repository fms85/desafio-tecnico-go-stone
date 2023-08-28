sync:
	@go mod tidy

start:
	@go run cmd/app/main.go

test:
	@go test ./... -timeout 5s -cover -coverprofile=cover.out
	@go tool cover -html=cover.out -o cover.html

sec:
	@if [ ! -f "$(GOPATH)/bin/gosec" ]; then \
		echo "Gosec not found. Installing Gosec..."; \
		go install github.com/securego/gosec/v2/cmd/gosec@latest; \
	fi
	@gosec ./...