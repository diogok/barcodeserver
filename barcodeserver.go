package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/codabar"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/qr"
	"github.com/boombuler/barcode/twooffive"
	"github.com/julienschmidt/httprouter"
)

func Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bartype := p.ByName("type")
	content := r.URL.Query().Get("content")

	swidth := r.URL.Query().Get("width")
	sheight := r.URL.Query().Get("height")

	if swidth == "" {
		swidth = "0"
	}
	if sheight == "" {
		sheight = "0"
	}

	width, werr := strconv.Atoi(swidth)
	if werr != nil {
		log.Println(werr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", werr)
		return
	}
	height, herr := strconv.Atoi(sheight)
	if herr != nil {
		log.Println(herr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", herr)
		return
	}

	var err error
	var initialBarcode barcode.Barcode

	switch bartype {
	case "datamatrix":
		initialBarcode, err = datamatrix.Encode(content)
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 256
		}
	case "qr":
		initialBarcode, err = qr.Encode(content, qr.Q, qr.Auto)
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 256
		}
	case "codabar":
		initialBarcode, err = codabar.Encode(content)
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 50
		}
	case "code128":
		initialBarcode, err = code128.Encode(content)
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 25
		}
	case "code39":
		initialBarcode, err = code39.Encode(content, true, true)
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 25
		}
	case "ean":
		initialBarcode, err = ean.Encode(content)
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 25
		}
	case "2of5":
		initialBarcode, err = twooffive.Encode(content, true)
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 25
		}
	case "twooffive":
		initialBarcode, err = twooffive.Encode(content, true)
		if width == 0 {
			width = 256
		}
		if height == 0 {
			height = 25
		}
	}

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	} else if bartype == "" || initialBarcode == nil {
		log.Println("Bad type")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", "Bad bar type")
		return
	} else {
		var serr error
		var finalBarcode barcode.Barcode
		if width != 0 && height != 0 {
			finalBarcode, serr = barcode.Scale(initialBarcode, width, height)
		} else {
			finalBarcode = initialBarcode
		}
		if serr != nil {
			log.Println(serr)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", serr)
			return
		}

		imageBuffer := new(bytes.Buffer)
		perr := png.Encode(imageBuffer, finalBarcode)
		if perr != nil {
			log.Println(perr)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", perr)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", strconv.Itoa(len(imageBuffer.Bytes())))

		if _, err := w.Write(imageBuffer.Bytes()); err != nil {
			log.Println("unable to write image.")
		}
	}
}

func main() {
	sport := os.Getenv("PORT")
	var port string
	if len(sport) > 2 {
		port = sport
	} else {
		port = "8080"
	}

	log.Printf("Binding at 0.0.0.0:%s", port)

	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public"))

	router.GET("/", Get)
	router.GET("/:type", Get)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
