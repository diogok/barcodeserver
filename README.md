# BarCodeServer

Simple HTTP API around [https://github.com/boombuler/barcode](https://github.com/boombuler/barcode).

## Api Usage

Simply issue a get using the type of barcode desired as path and the content as queryString content and it will return a redirect to generated png.

  $ curl -X GET http://localhost:8080/datamatrix?content=Whatever%20data

Available types:

* codabar
* code128
* code39
* ean
* datamatrix
* qr
* 2of5

In case of error it will return Bad Request 400 and error in plain/text.

## Deploy

Using the binary, download from the release page and run the binary and it will bind to port 8080.

Using docker:

  $ docker run -p 8080:8080 -v /tmp/barcodes:/opt/barcode/public diogok/barcodeserver


It will save generated code at public folder. You can delete generated artefacts and if requested server will regenerate.


## License 

MIT , same as [https://github.com/boombuler/barcode](https://github.com/boombuler/barcode).

