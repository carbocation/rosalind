package main

import (
    "fmt"
    "github.com/carbocation/rosalind.git/rosalind"
)

func main() {
	in := []byte("AAAACCCGGT")
    out := rosalind.ReverseComplement(in)
    
    fmt.Println(string(out))
}