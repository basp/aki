// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a modern LambdaMOO clone.
package main

import (
    "testing"
)

func TestActivation(t *testing.T) {
    a := new(Activation)
    a.Push(Var{Num:123})
    a.Push(Var{Str:"foo"})
    var v Var
    v, _ = a.Pop()
    if v.Str != "foo" {
        t.Error("Expected foo but got", v.Str)
    }
    v, _ = a.Pop()
    if v.Num != 123 {
        t.Error("Expected 123 but got", v.Num)
    }
}

func TestPopEmpty(t *testing.T) {
    vm := new(VM)
    a, ok := vm.Pop()
    if a != nil {
        t.Error("Expected nil but got", a)
    }
    if ok {
        t.Error("Expected false but got", ok)
    }
}

func setupForTwoPlusThree() *Program{
    code := []byte {
        IMM,
        0, 0, 0, 0,
        IMM,
        0, 0, 0, 1,
        ADD,
        RET,
    }
    literals := []Var {
        Var{Num:2},
        Var{Num:3},
    }
    return &Program {
        Main: code,
        Literals: literals,
    }
}

func TestExecute(t *testing.T) {
    prog := setupForTwoPlusThree()
    r := run(NewVM(prog, -1))
    if r.Num != 5 {
        t.Error("Expected 5 but got", r.Num)
    }
}

func BenchmarkTwoPlusThree(b *testing.B) {
    for i := 0; i < b.N; i++ {
        prog := setupForTwoPlusThree()
        Execute(prog)
    }
}

func TestVM(t *testing.T) {
    vm := new(VM)
    a1 := new(Activation)
    a2 := new(Activation)
    vm.Push(a1)
    vm.Push(a2)
    var a *Activation
    a, _ = vm.Pop()
    if a != a2 {
        t.Error("Expected", a2, "but got", a)
    }
    a, _ = vm.Pop()
    if a != a1 {
        t.Error("Expected", a1, "but got", a)
    }
}