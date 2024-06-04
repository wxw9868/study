package main

import (
	"bytes"
	"image/png"
	"net/http"

	"github.com/aofei/cameron"
)

func main() {
	http.ListenAndServe("localhost:8080", http.HandlerFunc(identicon))
}

func identicon(rw http.ResponseWriter, req *http.Request) {
	buf := bytes.Buffer{}
	png.Encode(&buf, cameron.Identicon([]byte(req.RequestURI), 540, 60))
	rw.Header().Set("Content-Type", "image/png")
	buf.WriteTo(rw)
}
