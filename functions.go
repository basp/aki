// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a LambdaMOO clone.
package main

type PackageKind int

const (
    BI_RETURN PackageKind = iota
    BI_RAISE
    BI_CALL
    BI_SUSPEND
)

type Package struct {
    Kind PackageKind
    Ret Var
    Raise struct {
        Code Var
        Msg string
        Value Var        
    }
    Call struct {
        PC int
        Data interface{}        
    }
    Suspend struct {
        Proc func(*VM, interface{}) Error
        Data interface{}        
    }
}