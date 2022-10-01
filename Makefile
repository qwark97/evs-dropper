run:
	@go run ./cmd/ev-dropper-server/main.go

build:
	go build ./cmd/ev-dropper-server/

unit:
	go test ./pkg/...

test: unit
