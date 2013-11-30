// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a LambdaMOO clone.
package main

const (
    NOP byte = iota
    IF
    WHILE
    EIF
    FORK
    FORK_WITH_ID
    FOR_LIST
    FOR_RANGE
    CALL_VERB
    BI_FUNC_CALL
    MAKE_SINGLETON_LIST
    MUL
    DIV
    MOD
    ADD
    SUB
    EQ
    NE
    LE
    GT
    GE
    IN
    AND
    OR
    NEG
    NOT
    IMM
    MAKE_EMPTY_LIST
    LIST_ADD_TAIL
    LIST_APPEND
    RETURN
    RETURN0
    POP
    EXT
)

const (
    OptinumStart = int(EXT) + 1
    MaxOpcode = 255
    OptinumLo = -10
    OptinumHi = MaxOpcode - OptinumStart + OptinumLo
)

func IsOptinumOpcode(o byte) bool {
    return int(o) >= OptinumStart
}

func OpcodeToOptinum(o byte) int {
    return int(o) - OptinumStart + OptinumLo
}

func InOptinumRange(i int) bool {
    return i >= OptinumLo && i <= OptinumHi
}

func OptinumToOpcode(i int) byte {
    return byte(OptinumStart + i - OptinumLo)
}