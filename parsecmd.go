package main

import (
    "strings"
)

type parsedCmd struct {
    verb string
    argstr string
    args Val
    dobjstr string
    dobj Objid
    prepstr string
    iobjstr string
    iobj Objid
}

func tokenize(s string) []string {
    words := make([]string, 0, 32)
    word := ""
    ignoringSpaces := false
    for _, c := range s {
        switch {
            case c == ' ':
                if !ignoringSpaces {
                    words = append(words, word)
                    word = ""
                } else {
                    word += string(c)
                }
                break
            case c == '"':
                ignoringSpaces = !ignoringSpaces
                break
            default:
                word += string(c)
        }
    }
    if len(word) > 0 {
        words = append(words, word)
    }
    return words
}

var prepositions = []string {
    "with",
    "using",
    "at",
    "to",
    "in front of",
    "in",
    "inside",
    "into",
    "on top of",
    "on",
    "onto",
    "upon",
    "out of",
    "from inside",
    "from",
    "over",
    "through",
    "under",
    "underneath",
    "beneath",
    "behind",
    "for",
    "about",
    "is",
    "as",
    "off",
    "off of",
}

func findPreposition(tokens []string) (start int, end int, ok bool) {
    var maybe string
    for _, p := range prepositions {
        for i, _ := range tokens {
            if p == tokens[i] {
                start, end, ok = i, i, true
                return
            }
            if i < len(tokens) - 1 {
                maybe = strings.Join(tokens[i:i + 1], " ")
                if p == maybe {
                    start, end, ok = i, i + 1, true
                    return
                }
            }
            if i < len(tokens) - 2 {
                maybe = strings.Join(tokens[i:i + 2], " ")
                if p == maybe {
                    start, end, ok = i, i + 2, true
                    return
                }
            }
        }
    }
    return
}

func parse(tokens []string) (verb string, dobj string, iobj string, ok bool) {
    if len(tokens) == 0 {
        return
    }
    ok = true
    verb = tokens[0]
    start, end, ok := findPreposition(tokens)
    if ok {
        dobj = strings.Join(tokens[1:start], " ")
        iobj = strings.Join(tokens[end + 1:], " ")
    } else {
        dobj = strings.Join(tokens[1:], " ")
    }
    return
}