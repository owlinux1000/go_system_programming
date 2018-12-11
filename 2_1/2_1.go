package main

import (
    "os"
    "fmt"
)

func main() {

    fmt.Fprintf(os.Stdout, "%d %f %s", 10, 0.349, "hogehoge")
    
}
