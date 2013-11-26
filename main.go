package main

import (
    "fmt"
    "log"
)

const (
    SystemObject = 0
    Nothing = -1
    Ambiguous = -2
    FailedMatch = -3   
)

type Error int

const (
    E_NONE Error = iota
    E_TYPE
    E_DIV
    E_PERM
    E_PROPNF
    E_VERBNF
    E_VARNF
)

type Type int

const (
    TYPE_NONE Type = iota
    TYPE_INT
    TYPE_FLOAT
    TYPE_STR
    TYPE_OBJ
    TYPE_LIST
    TYPE_ERR
)

type Objid int

func (o Objid) String() string {
    return fmt.Sprintf("#%v", int(o))
}

type Val struct {
    Num int
    Fnum float32
    Str string
    Obj Objid
    List []Val
    Err Error
    Type Type
}

func createObjects() {
    for i := 0; i < 100; i++ {
        o := newObject()
        log.Printf("Created %v", o)        
    }
}

func main() {
    go createObjects()
    go createObjects()
    go createObjects()
    _, _ = fmt.Scanln()
}