.PHONY: build clean deploy

build:
	# dep ensure -v
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/main internal/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --stage dev --verbose

test:
	go test -v ./internal/...