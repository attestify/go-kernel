test:
	go test -race -cover -coverprofile=cover.out ./... -count=1

cover:
	go tool cover -html=cover.out

fmt:
	go fmt ./...