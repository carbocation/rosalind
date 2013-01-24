package rosalind

import (
    "strings"
)

//#1 Given a string, counts the bases in that string
//Does no checking to make sure that the string contains valid DNA letters
func CountBases(Dna []byte) map[string]int {
    out := map[string]int{"A": 0, "C": 0, "T": 0, "G": 0}

    for _, base := range Dna {
        out[string(base)]++
    }

    return out
}

//#2 Given a byte array, returns a copy having all "T" replaced with "U"
func TranscribeDna(Dna []byte) []byte {
    Rna := make([]byte, len(Dna))
    copy(Rna, Dna)

    for i, base := range Rna {
        if string(base) == "T"{
            Rna[i] = 'U'
        }
    }

    return Rna
}

func ReverseComplement(Dna []byte) []byte {
    revc := reverseBytes(Dna)

    for i, base := range revc {
        revc[i] = complementDna(base)
    }
    
    return revc
}

//Given a FASTA sequence of the following structure, return a hashmap
// with the names of the sequences as the keys, and the sequences with 
// linebreaks removed as the values
//Example:
//>RedRobin
//GATTACAGATTACAGATTACAGATTACAGATTACA
//GATTACA
// yields
//map[RedRobin: GATTACAGATTACAGATTACAGATTACAGATTACAGATTACA]
func ParseFasta(in string) map[string]string {
    entries := map[string]string{}

    var name string
    
    //We consider that there may be more than one FASTA sequence in this string
    for _, line := range strings.Split(in, "\n") {
        if line[0:1] == ">" {
            name = line[1:]
            entries[name] = ""
            continue
        }

        entries[name] += strings.Trim(line, "\n")
    }

    return entries
}

//Calculate the GC fraction [0-1] of a Dna byte array
func FracGc(Dna string) float64{
    var bases, gcBases int

    bases = len(Dna)
    gcBases = 0.0

    for _, base := range Dna{
        if base == 'G' || base == 'C' {
            gcBases += 1
        }
    }

    return float64(gcBases)/float64(bases)
}