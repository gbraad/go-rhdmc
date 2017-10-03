all:
	#go get -v -d ./...
	go test
	go install -v ./cmd/rhdmc
