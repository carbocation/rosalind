package rosalind

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

//Non-exported functions, i.e., functions that simply support the primary functions in this package

func complementDna(base byte) byte {
    switch{
        case base == 'A':
            return 'T'
        case base == 'T':
            return 'A'
        case base == 'C':
            return 'G'
        case base == 'G':
            return 'C'
    }

    return 0
}

func reverseBytes(a []byte) []byte {
    b := make([]byte, len(a))
    j := len(b) - 1

    for i := 0; i <= j; i++ {
        b[j-i] = a[i]
    }

    return b
}