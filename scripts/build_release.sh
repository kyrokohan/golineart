# Linux
CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o dist/golineart-linux-amd64 ./cmd/gla/main.go
CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build -o dist/golineart-linux-arm64 ./cmd/gla/main.go
# macOS
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o dist/golineart-darwin-amd64 ./cmd/gla/main.go
CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o dist/golineart-darwin-arm64 ./cmd/gla/main.go
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/golineart-windows-amd64.exe ./cmd/gla/main.go
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o dist/golineart-windows-arm64.exe ./cmd/gla/main.go
