package main

import (
    "os"
    "fmt"
    "time"
    "strconv"
)


func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: 4_1 <SECOND>")
        os.Exit(0)
    }
    i, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    limit := time.After(time.Second * time.Duration(i))
    for {
        select {
        case <-limit:
            fmt.Printf("After %s second\n", os.Args[1])
            os.Exit(0)
        default:
            time.Sleep(1 * time.Second)
            fmt.Printf("Working now...\n")
        }
    }
    
}
