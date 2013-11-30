// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a LambdaMOO clone.
package main

import (
    "time"
    "log"
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

func (vm *VM) Push(a *Activation) {
    vm.Stack = append(vm.Stack, a)
}

func NewActivation(prog *Program, vector int) *Activation {
    return &Activation{Prog: prog, Vector: vector}
}

func NewVM(prog *Program, vector int) *VM {
    vm := new(VM)
    vm.Push(NewActivation(prog, vector))
    return vm
}

func Execute(prog *Program) {
    rt <- NewVM(prog, -1)
}

func Delay(prog *Program, seconds int) {
    exec(prog, seconds, -1)
}

var LOG = 0
var rt = make(chan *VM)

func toInt(bs []byte) int {
    i := int(bs[0]) << 24 
    i |= int(bs[1] << 16) 
    i |= int(bs[2]) << 8 
    i |= int(bs[3])
    return i
}

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

func exec(prog *Program, delay int, vector int) {
    vm := NewVM(prog, vector)
    suspend(vm, delay)
}

func suspend(vm *VM, seconds int) {
    go func() {
        if seconds > 0 {
            time.Sleep(time.Duration(seconds) * time.Second)
        }        
        rt <- vm
    }()
}

func run(vm *VM) Var {
    a, ok := vm.Pop()
    if !ok {
        log.Fatal("No activation frames")
    }
    var v []byte
    for {
        switch a.Vector {
        case MainVector: v = a.Prog.Main
        default: v = a.Prog.Forks[a.Vector]
        }
        op := v[a.PC]
        a.PC++
        switch op {
        case NOP:
            continue
        case POP:
            a.Pop()
            continue;
        case IMM:
            i := toInt(v[a.PC:a.PC + 4])
            a.PC += 4
            a.Push(a.Prog.Literals[i])
            continue
        case MAKE_EMPTY_LIST:
            a.Push(NewList([]Var { }))
            continue
        case ADD, SUB, MUL, DIV, MOD:
            v1, _ := a.Pop()
            v2, _ := a.Pop()
            var r Var
            switch op {
            case ADD: r = v1.Add(v2)
            case SUB: r = v1.Sub(v2)
            case MUL: r = v1.Mul(v2)
            case DIV: r = v1.Div(v2)
            case MOD: r = v1.Mod(v2)
            }
            a.Push(r)
            continue
        case RETURN, RETURN0:
            var r Var
            switch op {
            case RETURN: r, _ = a.Pop()
            case RETURN0: r = NewInt(0)
            }
            a, ok = vm.Pop()
            if ok {
                a.Push(r)
                continue
            }
            // We are done.
            return r
        case CALL_VERB:
            vm.Push(a)
            a = NewActivation(printFoo, MainVector)
            continue
        case MAKE_SINGLETON_LIST:
            v, _ := a.Pop()
            a.Push(NewList([]Var { v }))           
        case BI_FUNC_CALL:
            args, _ := a.Pop()
            if args.Type != LIST {
                return NewErr(E_INVARG)
            }
            i := toInt(v[a.PC:a.PC + 4])
            a.PC += 4
            r, suspended := bifuncs[i](args, a, vm)
            if suspended {
                return r
            }
            a.Push(r)
            continue
        case FORK:
            i := toInt(v[a.PC:a.PC + 4])
            a.PC += 4
            d := toInt(v[a.PC:a.PC + 4])
            a.PC += 4
            exec(a.Prog, d, i)
            continue
        default:
            if !IsOptinumOpcode(op) {
                log.Fatalf("Unknown opcode: %v", op)
            }
            r := NewInt(OpcodeToOptinum(op))
            a.Push(r)
            continue
        }
    }    
}

var printFoo = &Program{
    Main: []byte {
        IMM,
        0, 0, 0, 0,
        MAKE_SINGLETON_LIST,
        BI_FUNC_CALL,
        0, 0, 0, 1,
        RETURN0,
    },
    Literals: []Var {
        NewStr("foo"),
    },
}

func bf_suspend(args Var, a *Activation, vm *VM) (Var, bool) {
    if len(args.List) <= 0 {
        return NewErr(E_INVARG), false
    }
    if d := args.List[0]; d.Type != INT {
        return NewErr(E_INVARG), false
    } else {
        // We need to push back the current frame
        // because it will be consumed (again) by
        // the run function later.
        vm.Push(a)
        suspend(vm, d.Num)
        return d, true
    }
    return NewErr(E_NONE), false    
}

func bf_notify(args Var, a *Activation, vm *VM) (Var, bool) {
    if len(args.List) <= 0 {
        return NewErr(E_INVARG), false
    }
    if m := args.List[0]; m.Type != STR {
        return NewErr(E_INVARG), false
    } else {
        log.Println(m.Str)
        return m, false
    }
    return NewErr(E_NONE), false
}

var bifuncs = []func(Var, *Activation, *VM) (result Var, done bool) {
    bf_suspend,
    bf_notify,
}