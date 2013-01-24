/*
Takes in a mult-entry FASTA file and returns the name of the 
entry with the highest GC content (one one line), followed in the next line
with the GC % of that entry.
*/
package main

import (
    "fmt"
    "github.com/carbocation/rosalind.git/rosalind"
)

func main() {
    //Build a hashmap with fasta names as the keys, and the corresponding 
    //FASTA sequence (with linebreaks removed) as the values
    fasta := rosalind.ParseFasta(`>Rosalind_6404
CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC
TCCCACTAATAATTCTGAGG
>Rosalind_5959
CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT
ATATCCATTTGTCAGCAGACACGC
>Rosalind_0808
CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC
TGGGAACCTGCGGGCAGTAGGTGGAAT`)
    
    bestEntry := ""
    bestGc := 0.0

    //Iterate over each map entry; if it's got higher GC content than the last, 
    //that's our new winner.
    for name, seq := range fasta {
        fastaGc := rosalind.FracGc(seq)

        if fastaGc > bestGc {
            bestEntry = name
            bestGc = fastaGc
        }
    }

    //Print the output as demanded
    fmt.Printf("%s\n%f%%\n", bestEntry, 100.0*bestGc)
}