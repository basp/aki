package main

import (
    "testing"
)

func TestMatchObject(t *testing.T) {
    p := CreateObject()
    loc := CreateObject()
    o1 := CreateObject()
    o2 := CreateObject()
    o3 := CreateObject()
    o4 := CreateObject()    
    ChangeLocation(p, loc)
    ChangeLocation(o1, p)
    ChangeLocation(o2, p)
    ChangeLocation(o3, p)
    ChangeLocation(o4, loc)
    objects[o1].Name = "foo"
    objects[o2].Name = "foobar"
    objects[o3].Name = "fooboz"   
    objects[o4].Name = "quux" 
    var r Objid
    r = MatchObject(p, "")
    if r != Nothing {
        t.Error("Expected Nothing but got", r)
    }
    r = MatchObject(p, "#456")
    if r != Objid(456) {
        t.Error("Expected", Objid(456), "but got", r)
    }
    r = MatchObject(p, "me")
    if r != p {
        t.Error("Expected", p, "but got", r)
    }
    r = MatchObject(p, "here")
    if r != loc {
        t.Error("Expected", loc, "but got", r)
    }
    r = MatchObject(p, "foo")
    if r != o1 {
        t.Error("Expected", o1, "but got", r)
    }
    r = MatchObject(p, "foob")
    if r != Ambiguous {
        t.Error("Expected Ambiguous but got", r)
    }
    r = MatchObject(p, "quux")
    if r != o4 {
        t.Error("Expected", o4, "but got", r)
    }
}