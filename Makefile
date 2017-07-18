GOPATH=$(CURDIR)

build: *.go
	go build

docker-image: echo-linux
	docker build -t echo:standalone -f Dockerfile.standalone .
	docker build -t echo:alpine     -f Dockerfile.alpine     .

release:
	docker tag echo:standalone kesselborn/echo:standalone
	docker push kesselborn/echo:standalone

	docker tag echo:standalone kesselborn/echo:latest
	docker push kesselborn/echo:latest

	docker tag echo:alpine kesselborn/echo:alpine
	docker push kesselborn/echo:alpine

echo-linux: *.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o $@ .
