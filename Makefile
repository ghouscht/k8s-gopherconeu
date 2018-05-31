PROJECT?=github.com/ghouscht/k8s-gopherconeu
GOOS?=linux
APP?=k8s-gopherconeu
PORT?=8080

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

CONTAINER_IMAGE?=docker.io/thomasgosteli/gophercon

GOOS?=linux
GOARCH?=amd64

clean:
	rm -rf ./bin/${APP}

test:
	go test -race -cover ./...

build: clean
	CGO_ENABLED=0
	go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${GOOS}-${GOARCH}/${APP} ${PROJECT}/cmd 

run: build
	SERVICE_PORT=${PORT} ./bin/${APP}

push: build
	docker push $(CONTAINER_IMAGE):$(RELEASE)

deploy: push
	helm upgrade ${CONTAINER_NAME} -f gophercon/values.yaml charts --kube-context ${KUBE_CONTEXT} --namespace ${NAMESPACE} --version=${RELEASE} -i --wait
