package main

import (
    "io"
    "os"
    "encoding/csv"
)

func main() {
    
    file, err := os.Create("./2_2.csv")
    if err != nil {
        panic(err)
    }
    
    writer := io.MultiWriter(os.Stdout, file)
    csv_writer := csv.NewWriter(writer)
    csv_writer.Write([]string{
        "id",
        "name",
        "age",
    })
    csv_writer.Write([]string{
        "1",
        "owlinux1000",
        "22",
    })
    csv_writer.Flush()
}
