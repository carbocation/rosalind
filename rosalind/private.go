//Non-exported functions, i.e., functions that simply support the primary functions in this package
package rosalind

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