// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a modern LambdaMOO clone.
package main

import (
    "fmt"
)

func main() {   
    LOG = 1
    // Main vector that has two forks and a suspend
    main := []byte {
        IMM,
        0, 0, 0, 0,     // push slot 0 from literals
        IMM,
        0, 0, 0, 1,     // push slot 1 from literals
        ADD,            // pop lhs and rhs from stack and push lhs + rhs
        FORK,           
        0, 0, 0, 0,     // fork index 0
        0, 0, 0, 2,     // delay for 2 seconds
        FORK,
        0, 0, 0, 1,     // fork index 1
        0, 0, 0, 2,     
        FORK,
        0, 0, 0, 2,     // fork index 2
        0, 0, 0, 3,     // 1 second later as the other forks     
        IMM,            // Add the same slots again
        0, 0, 0, 0,     
        IMM,
        0, 0, 0, 1,     
        ADD,
        SUSPEND,        
        0, 0, 0, 5,     // suspend 5 seconds
        ADD,            
        RET,            // returns 2 * (slot 0 + slot 1)
    }
    // This fork adds two Int vars from the IMM slots.
    fork1 := []byte { 
        IMM,
        0, 0, 0, 3,    
        IMM,
        0, 0, 0, 4,    
        ADD,            
        RET,
    }
    // This fork adds two Float vars from the IMM slots.
    fork2 := []byte {
        IMM,
        0, 0, 0, 5,
        IMM,
        0, 0, 0, 6,
        ADD,
        RET,
    }
    // This fork adds two numbers in optinum range.
    fork3 := []byte {
        OptinumToOpcode(2),
        OptinumToOpcode(3),
        ADD,
        RET,
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
    }
    prog := &Program {
        Main: main,
        Forks: [][]byte { fork1, fork2, fork3 },
        Literals: literals,
    }
    Execute(prog)
    _, _ = fmt.Scanln()
}

// Output
// ------
// Suspend main:
//      2013/11/28 23:53:34 => {0 0  [] 0 0 0}
//
// Fork 0 and 1 finish at the same time:
//      2013/11/28 23:53:36 => {0 5  [] 0 0 1}
//      2013/11/28 23:53:36 => {0 0 foobar [] 0 0 2}
//
// Main finishes (after suspend):
//      2013/11/28 23:53:39 => {492 0  [] 0 0 0}