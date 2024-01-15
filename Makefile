GO_VERSION:= 1.20.2# Go version to use
setup:
	install-go init-go install-lint
build:
	go build -o api cmd/main.go
test:
	go test ./... -coverprofile=coverage.out
coverage:
	go tool cover -func coverage.out | grep "total:"  | \
	awk '{print  ((int($$3)> 80 )!=1) ? "Coverage is less than 80%": "Coverage is greater than 80%" }'
report:
	go tool cover -html=coverage.out -o cover.html
check-format:
	test -z $$(go fmt ./...)
install-lint:
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.41.1
install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
    echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
    echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc
copy-hooks:
	chmod +x scripts/hooks/*
	cp scripts/hooks/* .git/hooks/
