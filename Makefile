include .env

.PHONY: build clean deploy

build:
	# dep ensure -v
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/main function/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls config credentials --provider aws --key ${AWS_LAMBDA_KEY} --secret ${AWS_LAMBDA_SECRET}
	# sls deploy --stage dev --verbose

test:
	go test -v ./function/...

mocks:
	mockery --all --dir function/ --output function/mocks