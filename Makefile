project = barcodeserver

all: build

arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -tags netgo -ldflags '-w'

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w'

clean:
	rm petals

docker: build
	docker build -t diogok/$(project) .

docker-arm: arm
	docker build -t diogok/$(project):arm .

push:
	docker push diogok/$(project)


