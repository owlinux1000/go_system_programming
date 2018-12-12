package main

import (
    "strings"
    "io"
    "os"
)

var (
    computer = strings.NewReader("COMPUTER")
    system = strings.NewReader("SYSTEM")
    programming = strings.NewReader("PROGRAMMING")
)

func main() {

    var stream io.Reader
    
    char_a := io.NewSectionReader(programming, 5, 1)
    char_s := io.NewSectionReader(system, 0, 1)
    char_c := io.LimitReader(computer, 1)
    char_i1 := io.NewSectionReader(programming, 8, 1)
    char_i2 := io.NewSectionReader(programming, 8, 1)
    
    stream = io.MultiReader(char_a, char_s, char_c, char_i1, char_i2)
    io.Copy(os.Stdout, stream)
}
