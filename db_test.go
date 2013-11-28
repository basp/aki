// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Aki is a modern LambdaMOO clone.
package main

import (
    "testing"
)

func TestChangeLocation(t *testing.T) {
    o1 := CreateObject()
    o2 := CreateObject()
    var r Objid
    r = ObjectLocation(o1)
    if r != Nothing {
        t.Error("Expected Nothing but got", r)
    }
    ChangeLocation(o2, o1)
    r = ObjectLocation(o2)
    if r != o1 {
        t.Error("Expected", o1, "but got", r)
    }
    front := objects[o1].Contents.Front()
    if front.Value != o2 {
        t.Error("Expected", o2, "but got", front.Value)
    }
    ChangeLocation(o2, Nothing)
    count := objects[o1].Contents.Len()
    if count != 0 {
        t.Error("Expected 0 but got", count)
    }
}