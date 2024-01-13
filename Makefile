GO_VERSION:= 1.20	# Go version to use
setup:
	echo "Setting up..."
build:
	go build -o api cmd/main.go
