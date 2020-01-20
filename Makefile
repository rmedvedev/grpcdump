.PHONY: deps
deps:
	go get -d -v -t ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -race ./... -coverprofile=coverage.txt -covermode=atomic

.PHONY: build
build:
	CGO_ENABLED=0 go build cmd/grpcdump/main.go

.PHONY: release
release:
	@GO111MODULE=off go get github.com/goreleaser/goreleaser
	goreleaser --rm-dist

.PHONY: lint
lint:
	@go get golang.org/x/lint/golint
	golint ./...