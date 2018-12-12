package main

import (
    "os"
    "crypto/rand"
)

func main() {
    
    random := make([]byte, 1024)
    _, err := rand.Reader.Read(random)
    if err != nil {
        panic(err)
    }

    file, err := os.Create("./3_2.bin")
    if err != nil {
        panic(err)
    }

    file.Write(random)
    
}
