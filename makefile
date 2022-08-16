clean:
	go clean -testcache

test: clean
	go test -race -timeout 30s -v ./...