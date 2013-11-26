package main

import (
    "fmt"
    "testing"
)

type c struct {
    input string
    expected Objid
}

func TestMatchObject(t *testing.T) {
    p := Objid(123)
    cases := []c {
        c{"", Nothing},
        c{"#", FailedMatch},
        c{"#foo", FailedMatch},
        c{"#-123", -123},
        c{"#123", 123},
        c{"me", p},
        c{"ME", p},
    }
    for _, c := range cases {
        actual := matchObject(Objid(p), c.input)
        if actual != c.expected {
            msg := fmt.Sprintf("Expected %v but got %v", c.expected, actual)
            t.Error(msg)
        }
    }
}