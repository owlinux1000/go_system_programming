package main

import (
    "os"
    "io"
)

func CopyN(dest io.Writer, src io.Reader, length int) {
    lReader := io.LimitReader(src, int64(length))
    io.Copy(dest, lReader)
}

func main() {
    
    in, err := os.Open("/dev/urandom")
    if err != nil {
        panic(err)
    }

    CopyN(os.Stdout, in, 8)
    
}
