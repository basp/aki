package main

import (
    "time"
    "log"
)

const (
    NOP byte = iota
    IMM
    ADD
    FORK
    RET
)

type Program struct {
    Main []byte
    Forks [][]byte
    Literals []Var
    VarNames []string
}

type Activation struct {
    Vector int 
    Prog *Program
    Env []Var 
    Stack []Var
    PC int
}

const MainVector = -1

type VM struct {
    Stack []*Activation
}

func (a *Activation) Pop() (v Var, ok bool) {
    s := a.Stack
    if len(s) == 0 {
        return
    }
    v, a.Stack, ok = s[len(s) - 1], s[:len(s) - 1], true
    return 
}

func (a *Activation) Push(v Var) {
    a.Stack = append(a.Stack, v)
}

func (vm *VM) Pop() (a *Activation, ok bool) {
    s := vm.Stack
    if len(s) == 0 {
        return
    }
    a, vm.Stack, ok = s[len(s) - 1], s[:len(s) - 1], true
    return
}

func NewVM(prog *Program, vector int) *VM {
    vm := new(VM)
    vm.Push(&Activation{Prog: prog, Vector: vector})
    return vm
}

func (vm *VM) Push(a *Activation) {
    vm.Stack = append(vm.Stack, a)
}

func toInt(bs []byte) int {
    i := int(bs[0]) << 24 | int(bs[1] << 16) | int(bs[2]) << 8 | int(bs[3])
    return i
}

var rt = make(chan *VM)
var out = make(chan Var)

var LOG = 0

func init() {
    go func() {
        for {
            vm := <-rt
            r := run(vm)
            if LOG > 0 {
                log.Printf("=> %v", r)                
            }
        }
    }()
}

func Execute(prog *Program) {
    rt <- NewVM(prog, -1)
}

func Enqueue(prog *Program, delay int) {
    enqueue(prog, delay, -1)
}

func enqueue(prog *Program, delay int, vector int) {
    go func() {
        if delay > 0 {
            time.Sleep(time.Duration(delay) * time.Second)            
        }
        rt <- NewVM(prog, vector)
    }()    
}

func run(vm *VM) Var {
    a, ok := vm.Pop()
    if !ok {
        panic("VM needs at least one activation frame!")
    }
    var v []byte
    if a.Vector == MainVector {
        v = a.Prog.Main
    } else {
        v = a.Prog.Forks[a.Vector]
    }
    for {
        op := v[a.PC]
        a.PC++
        switch op {
        case NOP:
            continue
        case IMM:
            i := toInt(v[a.PC:a.PC + 4])
            a.PC += 4
            a.Push(a.Prog.Literals[i])
            continue
        case ADD:
            v1, _ := a.Pop()
            v2, _ := a.Pop()
            a.Push(v1.Add(v2))
            continue
        case RET:
            result, _ := a.Pop()
            a, ok := vm.Pop()
            if ok {
                a.Push(result)
                continue
            }
            return result
        case FORK:
            i := toInt(v[a.PC:a.PC + 4])
            a.PC += 4
            d := toInt(v[a.PC:a.PC + 4])
            a.PC += 4
            enqueue(a.Prog, d, i)
            continue
        }
    }    
}