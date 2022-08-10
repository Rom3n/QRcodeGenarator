package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"image/png"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generator/", viewQRcodeHandler)
	http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	p := Page{Title: "QR Code Generator"}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func viewQRcodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")

	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(w, qrCode)
}
