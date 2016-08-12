project = barcodeserver

all: build

clean: 
	rm barcodeserver-*

barcodeserver-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a -tags netgo -ldflags '-w' -o barcodeserver-arm

barcodeserver-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o barcodeserver-amd64

amd64: barcodeserver-amd64

arm: barcodeserver-arm

docker: amd64
	cp barcodeserver-amd64 barcodeserver
	docker build -t diogok/$(project) .
	rm barcodeserver

docker-arm: arm
	cp barcodeserver-arm barcodeserver
	docker build -t diogok/$(project):arm .
	rm barcodeserver

push:
	docker push diogok/$(project)


