package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/zip")

	writer := zip.NewWriter(w)
	defer writer.Close()

	zipper, err := writer.Create("./3_4.go")
	if err != nil {
		panic(err)
	}
	in, err := os.Open("./3_4.go")
	if err != nil {
		panic(err)
	}
	io.Copy(zipper, in)
	writer.Flush()
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
