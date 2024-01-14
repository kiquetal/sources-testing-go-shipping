GO_VERSION:= 1.20	# Go version to use
setup:
	echo "Setting up..."
build:
	go build -o api cmd/main.go
test:
	go test ./... -coverprofile=coverage.out
coverage:
	go tool cover -func coverage.out | grep "total:"  | \
	awk '{print  ((int($$3)> 80 )!=1) ? "Coverage is less than 80%": "Coverage is greater than 80%" }'
report:
	go tool cover -html=coverage.out -o cover.html
