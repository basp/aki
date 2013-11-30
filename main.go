// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a LambdaMOO clone.
package main

import (
    "fmt"
)

func main() {   
    LOG = 1
    // Main vector that has two forks and a suspend
    main := []Opcode {
        IMM,
        0, 0, 0, 0,             // push slot 0 from literals
        IMM,
        0, 0, 0, 1,             // push slot 1 from literals
        ADD,                    // pop lhs and rhs from stack and push lhs + rhs
        IMM,
        0, 0, 0, 7,             // OBJ
        IMM,
        0, 0, 0, 8,             // "print"
        IMM,
        0, 0, 0, 5,             // "bar"
        MAKE_SINGLETON_LIST,
        CALL_VERB,
        POP,
        IMM,
        0, 0, 0, 7,             // OBJ
        IMM,
        0, 0, 0, 8,             // "print"
        IMM,
        0, 0, 0, 6,             // "foo"
        MAKE_SINGLETON_LIST,
        CALL_VERB,
        POP,
        FORK,           
        0, 0, 0, 0,             // fork index 0
        0, 0, 0, 2,             // delay for 2 seconds
        FORK,
        0, 0, 0, 1,             // fork index 1
        0, 0, 0, 2,     
        FORK,
        0, 0, 0, 2,             // fork index 2
        0, 0, 0, 3,             // 1 second later as the other forks     
        IMM,                    // Add the same slots again
        0, 0, 0, 0,     
        IMM,
        0, 0, 0, 1,     
        ADD,
        OptinumToOpcode(5),     // suspend for 5 seconds
        MAKE_SINGLETON_LIST,    // make arglist with INT 5
        BI_FUNC_CALL,
        0, 0, 0, 0,             // builtin function at slot 0 (suspend)
        ADD,            
        RETURN,                    // returns 2 * (slot 0 + slot 1)
    }
    // This fork adds two Float vars from the IMM slots.
    fork0 := []Opcode { 
        IMM,
        0, 0, 0, 3,    
        IMM,
        0, 0, 0, 4,    
        ADD,
        RETURN,
    }
    // This fork adds two Str vars from the IMM slots.
    fork1 := []Opcode {
        IMM,
        0, 0, 0, 5,
        IMM,
        0, 0, 0, 6,
        ADD,
        MAKE_SINGLETON_LIST,
        BI_FUNC_CALL,
        0, 0, 0, 1,             // builtin function at slot 1 (notify)
        RETURN,
    }
    // This fork adds two numbers in optinum range.
    fork2 := []Opcode {
        OptinumToOpcode(2),
        OptinumToOpcode(3),
        ADD,
        RETURN,
    }
    // Load all the stuff that is used by the IMM instructions
    literals := []Var {
        NewInt(123),
        NewInt(123),
        NewInt(5),
        NewFloat(2.0),
        NewFloat(3.0),
        NewStr("bar"),
        NewStr("foo"),
        NewObj(0),
        NewStr("print"),
    }
    prog := &Program {
        Main: main,
        Forks: [][]Opcode { fork0, fork1, fork2 },
        Literals: literals,
    }
    Execute(prog)
    _, _ = fmt.Scanln()
}