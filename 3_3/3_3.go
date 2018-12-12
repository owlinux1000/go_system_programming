package main

import (
    "io"
    "os"
    "archive/zip"
)

func main() {

    out, err := os.Create("sample.zip")
    if err != nil {
        panic(err)
    }
    
    writer := zip.NewWriter(out)
    defer writer.Close()

    filename := os.Args[0]
    in, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    zipper, err := writer.Create(filename)
    if err != nil {
        panic(err)
    }
    
    io.Copy(zipper, in)
    
}
