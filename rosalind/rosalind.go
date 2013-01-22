package rosalind

//#1 Given a string, counts the bases in that string
//Does no checking to make sure that the string contains valid DNA letters
func CountBases(Dna string) map[string]int {
    runes := []rune(Dna)
    
    out := map[string]int{"A": 0, "C": 0, "T": 0, "G": 0}

    for _, base := range runes {
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