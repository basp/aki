// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a modern LambdaMOO clone.
package main

import (
    "testing"
)

func BenchmarkLongStrings(b *testing.B) {
    str1 := "abcdefghijklmnopqrstuvwxyz abcdefghijklmnopqrstuvwxyz abcdefghijklmnopqrstuvwxyz abcdefghijklmnopqrstuvwxyz"
    str2 := "abcdefghijklmnopqrstuvwxyz abcdefghijklmnopqrstuvwxyz abcdefghijklmnopqrstuvwxyz abcdefghijklmnopqrstuvwxyz"
    for i := 0; i < b.N; i++ {      
        compare(str1, str2)
    }
}

func BenchmarkShortStrings(b *testing.B) {
    str1 := "abcdefghijk"
    str2 := "abcdefghijk"
    for i := 0; i < b.N; i++ {
        compare(str1, str2)
    }
}

func BenchmarkDifferentLengthStrings(b *testing.B) {
    str1 := "abcdefg"
    str2 := "abc"
    for i := 0; i < b.N; i++ {
        compare(str1, str2)
    }
}

type cmp struct {
    str1 string
    str2 string
    expected int
}

var cases = []cmp {
    {"a", "a", 0},
    {"a", "A", 0},
    {"abcdefg", "ABCDEFG", 0},
    {"abc", "abcd", -1},
    {"abcd", "abc", 1},
    {"abc1", "abc1", 0},
    {"abc1", "abc2", -1},
    {"abc2", "abc1", 1},
    {"1234", "1234", 0},
    {"1231", "1232", -1},
    {"1232", "1231", 1},
}

func TestStringCompare(t *testing.T) {
    for _, c := range cases {
        actual := compare(c.str1, c.str2)
        if actual != c.expected {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}