// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a modern LambdaMOO clone.
package main

import (
    "fmt"
    "strings"
)

type Objid int

type VarType int

type ErrorType int

const (
    E_NONE ErrorType = iota
    E_TYPE
    E_DIV
    E_PERM
    E_PROPNF
    E_VERBNF
    E_VARNF
    E_INVIND
    E_RECMOVE
    E_MAXREC
    E_RANGE
    E_ARGS
    E_NACC
    E_INVARG
    E_QUOTA
    E_FLOAT
)

const (
    INT VarType = iota
    FLOAT
    OBJ
    STR
    ERR
    LIST
    CLEAR       // clear properties' value slot
    NONE        // uninitialized variables
    CATCH       // on-stack marker for an exception handler
    FINALLY     // on-stack marker for a try-finally clause
)

type Var struct {
    Num int
    Fnum float32
    Str string
    List []Var
    Obj Objid
    Err ErrorType
    Type VarType
}

var Zero Var
var True = NewInt(1)
var False = NewInt(0)

type binOp map[VarType]map[VarType]func(Var, Var) Var

var cmpOp = binOp {
    INT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            switch {
            case lhs.Num < rhs.Num:
                return NewInt(-1)
            case lhs.Num > rhs.Num:
                return NewInt(1)
            default:
                return NewInt(0)
            }
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            f := float32(lhs.Num)
            switch {
            case f < rhs.Fnum:
                return NewInt(-1)
            case f > rhs.Fnum:
                return NewInt(1)
            default:
                return NewInt(0)
            }
        },  
    },
    FLOAT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            f := float32(rhs.Num)
            switch {
            case lhs.Fnum < f:
                return NewInt(-1)
            case lhs.Fnum > f:
                return NewInt(1)
            default:
                return NewInt(0)
            }
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            switch {
            case lhs.Fnum < rhs.Fnum:
                return NewInt(-1)
            case lhs.Fnum > rhs.Fnum:
                return NewInt(1)
            default:
                return NewInt(0)
            }
        },
    },
    STR: map[VarType]func(Var, Var) Var {
        STR: func(lhs Var, rhs Var) Var {
            return NewInt(strcmp(lhs.Str, rhs.Str))
        },
    },
    OBJ: map[VarType]func(Var, Var) Var {
        OBJ: func(lhs Var, rhs Var) Var {
            switch {
            case lhs.Obj < rhs.Obj:
                return NewInt(-1)
            case lhs.Obj > rhs.Obj:
                return NewInt(1)
            default:
                return NewInt(0)
            }
        },
    },
    ERR: map[VarType]func(Var, Var) Var {
        ERR: func(lhs Var, rhs Var) Var {
            switch {
            case lhs.Err < rhs.Err:
                return NewInt(-1)
            case lhs.Err > rhs.Err:
                return NewInt(1)
            default:
                return NewInt(0)
            }
        },
    },
}

var addOp = binOp {
    INT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var { 
            return NewInt(lhs.Num + rhs.Num) 
        },
        FLOAT: func(lhs Var, rhs Var) Var { 
            return NewFloat(float32(lhs.Num) + rhs.Fnum) 
        },
        STR: func(lhs Var, rhs Var) Var {
            return NewStr(string(lhs.Num) + rhs.Str)
        },
    },
    FLOAT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewFloat(lhs.Fnum + float32(rhs.Num))
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            return NewFloat(lhs.Fnum + rhs.Fnum)
        },
        STR: func(lhs Var, rhs Var) Var {
            return NewStr(ftoa(lhs.Fnum) + rhs.Str)
        },
    },
    STR: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewStr(lhs.Str + string(rhs.Num))
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            return NewStr(lhs.Str + ftoa(rhs.Fnum))
        },
        STR: func(lhs Var, rhs Var) Var {
            return NewStr(lhs.Str + rhs.Str)
        },
    },
}

var subOp = binOp {
    INT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewInt(lhs.Num - rhs.Num)
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            return NewFloat(float32(lhs.Num) - rhs.Fnum)
        },
    },
    FLOAT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewFloat(lhs.Fnum - float32(rhs.Num))
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            return NewFloat(lhs.Fnum - rhs.Fnum)
        },
    },
}

var mulOp = binOp {
    INT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewInt(lhs.Num * rhs.Num)
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            return NewFloat(float32(lhs.Num) * rhs.Fnum)
        },
        STR: func(lhs Var, rhs Var) Var {
            return NewStr(strings.Repeat(rhs.Str, lhs.Num))
        },
    },
    FLOAT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewFloat(lhs.Fnum * float32(rhs.Num))
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            return NewFloat(lhs.Fnum * rhs.Fnum)
        },
    },
    STR: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewStr(strings.Repeat(lhs.Str, rhs.Num))
        },
    },
}

var divOp = binOp {
    INT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewFloat(float32(lhs.Num) / float32(rhs.Num))
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            return NewFloat(float32(lhs.Num) / rhs.Fnum)
        },
    },
    FLOAT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewFloat(lhs.Fnum / float32(rhs.Num))
        },
        FLOAT: func(lhs Var, rhs Var) Var {
            return NewFloat(lhs.Fnum / rhs.Fnum)
        },
    },
}

var modOp = binOp {
    INT: map[VarType]func(Var, Var) Var {
        INT: func(lhs Var, rhs Var) Var {
            return NewInt(lhs.Num % rhs.Num)
        },
    },
}

func evalBinOp(op binOp, lhs Var, rhs Var) Var {
    if _, ok := op[lhs.Type]; !ok {
        return NewErr(E_TYPE)
    }
    if f, ok := op[lhs.Type][rhs.Type]; !ok {
        return NewErr(E_TYPE)
    } else {
        return f(lhs, rhs)
    }
}

func NewInt(v int) Var {
    return Var{Num: v, Type: INT}
}

func NewFloat(v float32) Var {
    return Var{Fnum: v, Type: FLOAT}
}

func NewStr(v string) Var {
    return Var{Str: v, Type: STR}
}

func NewList(v []Var) Var {
    return Var{List: v, Type: LIST}
}

func NewObj(v Objid) Var {
    return Var{Obj: v, Type: OBJ}
}

func NewErr(v ErrorType) Var {
    return Var{Err: v, Type: ERR}
}

func (v Var) Cmp(other Var) Var {
    if v.Type == LIST && other.Type == LIST {
        switch {
        case len(v.List) < len(other.List):
            return NewInt(-1)
        case len(v.List) > len(other.List):
            return NewInt(1)
        }
        for i := 0; i < len(v.List); i++ {
            x, y := v.List[i], other.List[i]
            if r := evalBinOp(cmpOp, x, y); r.Type != ERR && r.Num == 0 {
                continue
            } else {
                return r
            }
        }
        return NewInt(0)
    }
    return evalBinOp(cmpOp, v, other)        
}

func (v Var) Lt(other Var) Var {
    r := v.Cmp(other)
    switch {
    case r.Type == ERR:
        return r
    case r.Num < 0:
        return True
    default:
        return False
    }
}

func (v Var) Gt(other Var) Var {
    r := v.Cmp(other)
    switch {
    case r.Type == ERR:
        return r
    case r.Num > 0:
        return True
    default:
        return False
    }
}

func (v Var) Le(other Var) Var {
    r := v.Cmp(other)
    switch {
    case r.Type == ERR:
        return r
    case r.Num < 0 || r.Num == 0:
        return True
    default:
        return False    
    }
}

func (v Var) Ge(other Var) Var {
    r := v.Cmp(other)
    switch {
    case r.Type == ERR:
        return r
    case r.Num > 0 || r.Num == 0:
        return True
    default:
        return False
    }
}

func (v Var) Eq(other Var) Var {
    r := v.Cmp(other)
    switch {
    case r.Type == ERR:
        return r
    case r.Num == 0:
        return True
    default:
        return False
    }
}

func (v Var) In(other Var) Var {
    if other.Type != LIST {
        return NewErr(E_TYPE)
    }
    for i := 0; i < len(other.List); i++ {
        if v.Eq(other.List[i]).IsTrue() {
            return NewInt(1)
        }
    }
    return NewInt(0)
}

func (v Var) Add(other Var) Var {
    return evalBinOp(addOp, v, other)
}

func (v Var) Sub(other Var) Var {
    return evalBinOp(subOp, v, other)
}

func (v Var) Mul(other Var) Var {
    return evalBinOp(mulOp, v, other)
}

func (v Var) Div(other Var) Var {
    return evalBinOp(divOp, v, other)
}

func (v Var) Mod(other Var) Var {
    return evalBinOp(modOp, v, other)
}

func (v Var) IsTrue() bool {
    switch v.Type {
    case INT:
        return v.Num != 0
    case FLOAT:
        return v.Num != 0.0
    case OBJ:
        return int(v.Obj) >= 0
    case STR:
        return len(v.Str) > 0
    case LIST:
        return len(v.List) > 0
    default:
        return false
    }
}

func (v Var) Or(other Var) Var {
    if v.IsTrue() {
        return NewInt(1)
    }
    if other.IsTrue() {
        return NewInt(1)
    }
    return NewInt(0)
}

func (v Var) And(other Var) Var {
    if !v.IsTrue() {
        return NewInt(0)
    }
    if !other.IsTrue() {
        return NewInt(0)
    }
    return NewInt(1)
}

func (e ErrorType) String() string {
    switch e {
        case E_NONE: 
            return "E_NONE"
        case E_TYPE:
            return "E_TYPE"
        case E_DIV:
            return "E_DIV"
        case E_PERM:
            return "E_PERM"
        case E_PROPNF:
            return "E_PROPNF"
        case E_VERBNF:
            return "E_VERBNF"
        case E_VARNF:
            return "E_VARNF"
        case E_INVIND:
            return "E_INVIND"
        case E_RECMOVE:
            return "E_RECMOVE"
        case E_MAXREC:
            return "E_MAXREC"
        case E_RANGE:
            return "E_RANGE"
        case E_ARGS:
            return "E_ARGS"
        case E_NACC:
            return "E_NACC"
        case E_INVARG:
            return "E_INVARG"
        case E_QUOTA:
            return "E_QUOTA"
        case E_FLOAT:
            return "E_FLOAT"
        default:
            return string(e)
    }
}

func (t VarType) String() string {
    switch t {
    case INT:
        return "INT"
    case FLOAT:
        return "FLOAT"
    case OBJ:
        return "OBJ"
    case STR:
        return "STR"
    case ERR:
        return "ERR"
    case LIST:
        return "LIST"
    case CLEAR:
        return "CLEAR"
    case NONE:
        return "NONE"
    case CATCH:
        return "CATCH"
    case FINALLY:
        return "FINALLY"
    }
    return fmt.Sprintf("UNKNOWN(%v)", int(t))
}

func (oid Objid) String() string {
    return fmt.Sprintf("#%v", int(oid))
}