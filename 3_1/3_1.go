package main

import (
    "io"
    "os"
)

func main() {

    in, err := os.Open("old.txt")
    if err != nil {
        panic(err)
    }
    
    out, err := os.Create("new.txt")
    if err != nil {
        panic(err)
    }
    
    io.Copy(out, in)

}
