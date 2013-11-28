package main

import (
    "fmt"
)

func main() {   
    LOG = 1
    main := []byte {
        IMM,
        0, 0, 0, 0,
        IMM,
        0, 0, 0, 1,
        ADD,
        FORK,
        0, 0, 0, 0,
        0, 0, 0, 2,
        IMM,
        0, 0, 0, 0,
        IMM,
        0, 0, 0, 1,
        ADD,
        ADD,
        RET,
    }
    fork := []byte {
        IMM,
        0, 0, 0, 3,
        IMM,
        0, 0, 0, 4,
        ADD,
        RET,
    }
    literals := []Var {
        NewInt(123),
        NewInt(123),
        NewInt(5),
        NewFloat(2.0),
        NewFloat(3.0),
    }
    prog := &Program {
        Main: main,
        Forks: [][]byte { fork },
        Literals: literals,
    }
    Execute(prog)
    _, _ = fmt.Scanln()
}