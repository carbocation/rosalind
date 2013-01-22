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

func ParseFasta(in string) map[string]string {
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

    return entries
}