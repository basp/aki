package main

import (
    "fmt"
    "testing"
)

type a struct {
    input string
    expected []string
}

func TestTokenize(t *testing.T) {
    cases := []a {
        a{"foo bar quux", []string{"foo", "bar", "quux"}},
        a{`foo "bar quux" zoz`, []string{"foo", "bar quux", "zoz"}},
    }
    for _, c := range cases {
        actual := tokenize(c.input)
        if len(actual) != len(c.expected) {
            msg := fmt.Sprintf("Length was %v but expected %v", len(actual), len(c.expected))
            t.Error(msg)
        }
        for i, _ := range actual {
            if actual[i] != c.expected[i] {
                msg := fmt.Sprintf("Expected '%v' but got '%v'", c.expected, actual)
                t.Error(msg)
                return
            }
        }
    }
}

type b struct {
    input []string
    expected []string
}

func TestParse(t *testing.T) {
    cases := []b {
        b{[]string{"foo"}, []string{"foo", "", ""}},
        b{[]string{"foo", "bar"}, []string{"foo", "bar", ""}},
        b{[]string{"foo", "bar", "quux"}, []string{"foo", "bar quux", ""}},
        b{[]string{"foo", "bar", "quux", "in", "zoz"}, []string{"foo", "bar quux", "zoz"}},
    }
    for _, c := range cases {
        verb, dobj, iobj, _ := parse(c.input)
        actual := []string {verb, dobj, iobj}
        for i := 0; i < 3; i++ {
            if actual[i] != c.expected[i] {
                msg := fmt.Sprintf("Expected '%v' but got '%v'", c.expected[i], actual[i])    
                t.Error(msg)
            }
        }
    }      
}