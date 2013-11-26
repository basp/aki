package main

import (
    "fmt"
    "testing"
)

func init() {
    objects[Objid(1)] = new(object)
}

type d struct {
    input Objid
    expected bool
}

func TestValid(t *testing.T) {
    cases := []d {
        d{Objid(123), false},
        d{Objid(1), true},
    }
    for _, c := range cases {
        actual := valid(c.input)
        if actual != c.expected {
            msg := fmt.Sprintf("Expected %v but got %v", c.expected, actual)
            t.Error(msg)
        }
    }
}