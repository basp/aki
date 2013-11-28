package main

func hasPrefix(s, prefix string) bool {
    if len(prefix) > len(s) {
        return false
    }
    for i := 0; i < len(prefix); i++ {
        c1 := s[i]
        c2 := prefix[i]
        if c1 >= 'A' && c1 <= 'Z' {
            c1 += 'a' - 'A'
        }
        if c2 >= 'A' && c2 <= 'Z' {
            c2 += 'a' - 'A'
        }
        if c1 != c2 {
            return false
        }
    }
    return true
}

func compare(str1 string, str2 string) int {
    lenstr1, lenstr2 := len(str1), len(str2)
    switch {
    case lenstr1 > lenstr2:
        return 1
    case lenstr2 > lenstr1:
        return -1
    }
    for i := 0; i < lenstr1; i++ {
        c1 := str1[i]
        c2 := str2[i]
        if c1 >= 'A' && c1 <= 'Z' {
            c1 += 'a' - 'A'
        }
        if c2 >= 'A' && c2 <= 'Z' {
            c2 += 'a' - 'A'
        }
        switch {
        case c1 > c2:
            return 1
        case c2 > c1:
            return -1
        }
    }
    return 0
}