package main

import (
    "fmt"
    "github.com/carbocation/rosalind.git/rosalind"
)

func main() {
	in := []byte("GATGGAACTTGACTACGTAAATT")
    out := rosalind.TranscribeDna(in)
    
    fmt.Println(string(out))
}