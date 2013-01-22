package main

import (
    "fmt"
    "github.com/carbocation/rosalind.git/rosalind"
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
    /*
    entries := map[string]string{}

    var name string

    for _, line := range strings.Split(in, "\n") {
        if line[0:1] == ">" {
            name = line[1:]
            entries[name] = ""
            continue
        }

        entries[name] += strings.Trim(line, "\n")
    }
    */

    fmt.Println(rosalind.ParseFasta(in))
}