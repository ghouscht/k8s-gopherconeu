PROJECT?=github.com/ghouscht/k8s-gopherconeu
GOOS?=linux
APP?=gophercon
PORT?=8080

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -rf ./dist/${APP}

test:
	go test -race -cover ./...

build: clean
	CGO_ENABLED=0
	go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./dist/${APP} ./cmd/gophercon.go

run: build
	SERVICE_PORT=${PORT} ./dist/${APP}
