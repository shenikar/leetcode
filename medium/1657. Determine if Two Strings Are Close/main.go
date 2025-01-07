package main

func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    m1 := make(map[byte]int)
    m2 := make(map[byte]int)
    for i := 0; i < len(word1); i++ {
        m1[word1[i]]++
        m2[word2[i]]++
    }

    if len(m1) != len(m2) {
        return false
    }

    for k, v := range m1 {
        if m2[k] == 0 {
            return false
        }
        delete(m2, k)
        if m2[k] != v {
            return false
        }
    }

    return true
}
