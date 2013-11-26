package main

import (
    "strconv"
)

func toLower(s string) string {
    b := make([]byte, len(s))
    for i := range b {
        c := s[i]
        if c >= 'A' && c <= 'Z' {
            c += 'a' - 'A'
        }
        b[i] = c
    }
    return string(b)
}

func matchObject(p Objid, name string) Objid {
    if len(name) == 0 {
        return Nothing
    }
    if name[0] == '#' {
        if num, err := strconv.Atoi(name[1:]); err != nil {
            return FailedMatch
        } else {
            return Objid(num)
        }
    }
    name = toLower(name)
    if name == "me" {
        return p
    }
    return Nothing
}