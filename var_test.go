// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a LambdaMOO clone.
package main

import (
    "testing"
)

type e struct {
    lhs Var
    rhs Var
    expected Var
}

func TestAdd(t *testing.T) {
    cases := []e {
        e{NewInt(2), NewInt(3), NewInt(5)},
        e{NewFloat(1.5), NewFloat(0.75), NewFloat(2.25)},
        e{NewStr("foo"), NewStr("bar"), NewStr("foobar")},
        e{NewStr("foo"), NewFloat(0.75), NewStr("foo0.75")},
    }
    for _, c := range cases {
        if actual := c.lhs.Add(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestSub(t *testing.T) {
    cases := []e {
        e{NewInt(2), NewInt(3), NewInt(-1)},
        e{NewFloat(1.5), NewFloat(0.75), NewFloat(0.75)},
        e{NewInt(2), NewFloat(1.5), NewFloat(0.5)},
    }
    for _, c := range cases {
        if actual := c.lhs.Sub(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestMul(t *testing.T) {
    cases := []e {
        e{NewInt(2), NewInt(3), NewInt(6)},
        e{NewFloat(0.5), NewInt(3), NewFloat(1.5)},
        e{NewInt(3), NewStr("foo"), NewStr("foofoofoo")},
        e{NewStr("foo"), NewInt(3), NewStr("foofoofoo")},
    }
    for _, c := range cases {
        if actual := c.lhs.Mul(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestDiv(t *testing.T) {
    cases := []e {
        e{NewInt(1), NewInt(2), NewFloat(0.5)},
    }
    for _, c := range cases {
        if actual := c.lhs.Div(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestMod(t *testing.T) {
    cases := []e {
        e{NewInt(3), NewInt(2), NewInt(1)},
    }
    for _, c := range cases {
        if actual := c.lhs.Mod(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestCmp(t *testing.T) {
    cases := []e {
        e{
            lhs: NewList([]Var { NewInt(1), NewInt(2), NewInt(3) }),
            rhs: NewList([]Var { NewInt(1), NewInt(2), NewInt(3) }),
            expected: NewInt(0),
        },
    }
    for _, c := range cases {
        if actual := c.lhs.Cmp(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestEq(t *testing.T) {
    cases := []e {
        e {
            lhs: NewList([]Var { NewInt(1), NewInt(2), NewInt(3) }),
            rhs: NewList([]Var { NewInt(1), NewInt(2), NewInt(3) }),
            expected: NewInt(1),
        },
        e {
            lhs: NewList([]Var { }),
            rhs: NewList([]Var { NewInt(1) }),
            expected: NewInt(0),
        },
    }
    for _, c := range cases {
        if actual := c.lhs.Eq(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestLt(t *testing.T) {
    cases := []e {
        e {
            lhs: NewList([]Var { NewInt(1), NewInt(2) }),
            rhs: NewList([]Var { NewInt(1), NewInt(3) }),
            expected: NewInt(1),
        },
    }
    for _, c := range cases {
        if actual := c.lhs.Lt(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestIsTrue(t *testing.T) {
    if NewInt(0).IsTrue() {
        t.Error("Expected false but got true")
    }
    if !NewInt(1).IsTrue() {
        t.Error("Expected true but got false")
    }
}

func TestOr(t *testing.T) {
    cases := []e {
        e{NewInt(1), NewInt(0), NewInt(1)},
        e{NewInt(0), NewInt(1), NewInt(1)},
        e{NewInt(0), NewInt(0), NewInt(0)},
        e{NewInt(1), NewInt(1), NewInt(1)},
    }
    for _, c := range cases {
        if actual := c.lhs.Or(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestAnd(t *testing.T) {
    cases := []e {
        e{NewInt(1), NewInt(0), NewInt(0)},
        e{NewInt(0), NewInt(1), NewInt(0)},
        e{NewInt(0), NewInt(0), NewInt(0)},
        e{NewInt(1), NewInt(1), NewInt(1)},
    }
    for _, c := range cases {
        if actual := c.lhs.And(c.rhs); !actual.Eq(c.expected).IsTrue() {
            t.Error("Expected", c.expected, "but got", actual)
        }
    }
}

func TestStringer(t *testing.T) {
    var s, expected string
    expected = "123"
    if s = NewInt(123).String(); s != expected {
        t.Error("Expected 123 but got", s)
    }
    expected = "#123"
    if s = NewObj(123).String(); s != expected {
        t.Error("Expected", expected, "but got", s)
    }
    expected = "[123, 'foo']"
    if s = NewList([]Var { NewInt(123), NewStr("foo") }).String(); s != expected {
        t.Error("Expected", expected, "but got", s)
    }
    expected = "123.45"
    if s = NewFloat(123.45).String(); s != expected {
        t.Error("Expected", expected, "but got", s)
    }
}