package main

import (
    "fmt"
    "github.com/carbocation/rosalind.git/rosalind"
    "strings"
    //"github.com/carbocation/util.git/functors"
    //"github.com/carbocation/util.git/datatypes"
)

func main() {
    in := `>Rosalind_6404
CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC
TCCCACTAATAATTCTGAGG
>Rosalind_5959
CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT
ATATCCATTTGTCAGCAGACACGC
>Rosalind_0808
CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC
TGGGAACCTGCGGGCAGTAGGTGGAAT`
    
    //Build a hashmap with fasta names as the keys, and the corresponding 
    //FASTA sequence (with linebreaks removed) as the values
    fasta := rosalind.ParseFasta(in)
    
    fmt.Println(fasta["Rosalind_0808"][0:1])
    
    bestFasta := ""
    bestGc := 0.0

    for name, seq := range fasta {
        fastaGc := rosalind.FracGc(seq)
        if fastaGc > bestGc {
            bestFasta = name
            bestGc = fastaGc
        }
    }

    fmt.Printf("%s\n%f%%", bestFasta, 100.0*bestGc)

    toLower := func(s string) interface{} {
        return strings.ToLower(s)
    }

    toByteArray := func(s string) interface{} {
        return []byte(s)
    }

    MapFuncOverStringStringHash := func(fx func(string) interface{}, mx map[string]string) map[string]interface{} {
        out := map[string]interface{}{}

        for i, val := range mx {
            out[i] = fx(val)
        }

        return out
    }

    //fmt.Println(MapFuncOverStringStringHash(f, fasta))
    fmt.Println(MapFuncOverStringStringHash(toLower, fasta))
    fmt.Println(MapFuncOverStringStringHash(toByteArray, fasta))
    
    //fmt.Printf("%f%%%s", 100.0*rosalind.FracGc([]byte(fmap["Rosalind_0808"])), "\n")
    //fmt.Printf("Hi %f%s", 100.0, "BRO")
}