package main

type Objid int

type VarType int

const (
    Int VarType = iota
    Float
    Str
    Obj
    List
    Err
)

type Var struct {
    Num int
    Fnum float32
    Str string
    List []Var
    Obj Objid
    Err int
    Type VarType
}

func NewInt(v int) Var {
    return Var{Num: v, Type: Int}
}

func NewFloat(v float32) Var {
    return Var{Fnum: v, Type: Float}
}

func NewStr(v string) Var {
    return Var{Str: v, Type: Str}
}

func NewList(v []Var) Var {
    return Var{List: v, Type: List}
}

func (v Var) Add(other Var) Var {
    switch v.Type {
    case Int:
        return NewInt(v.Num + other.Num)
    case Float:
        return NewFloat(v.Fnum + other.Fnum)
    case Str:
        return NewStr(v.Str + other.Str)
    case List:
        return NewList(append(v.List, other.List...))
    }
    return Var{Err:0}
}