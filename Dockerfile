FROM alpine:3.1

RUN mkdir -p /opt/barcode/public
VOLUME /opt/barcode/public

RUN apk add --update ca-certificates
COPY ./barcodeserver /opt/barcode/barcodeserver
CMD ["/opt/barcode/barcodeserver"]

