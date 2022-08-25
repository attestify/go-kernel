test:
	go test -race ./... -count=1

test-cover:
	go test -race -cover -coverprofile=cover.out ./... -count=1
	go tool cover -html=cover.out

fmt:
	go fmt ./...