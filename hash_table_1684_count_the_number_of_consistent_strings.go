func countConsistentStrings(allowed string, words []string) int {
    result := 0

    frequency := make(map[rune]bool)
    for _, allow := range allowed {
        frequency[allow] = true
    }

    for _, word := range words {
        consistent := true
        for _, w := range word {
           if !frequency[w] {
            consistent = false
            break
           }
        }

        if consistent {
            result++
        }
    }
    return result
}
