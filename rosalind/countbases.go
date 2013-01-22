package rosalind

//Given a string, counts the bases in that string
//Does no checking to make sure that the string contains valid DNA letters
func CountBases(Dna string) map[string]int {
    runes := []rune(Dna)
    
    out := map[string]int{"A": 0, "C": 0, "T": 0, "G": 0}

    for _, base := range runes {
        out[string(base)]++
    }

    return out
}