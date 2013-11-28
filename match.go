package main

import (
    "strconv"
)

const (
    SystemObject = 0
    Nothing = -1
    Ambiguous = -2
    FailedMatch = -3
)

func matchContents(p Objid, name string) (match Objid) {
    match = FailedMatch
    if !Valid(p) {
        return
    }
    loc := ObjectLocation(p)
    for _, id := range []Objid { p, loc } {
        if !Valid(id) {
            continue
        }
        contents := objects[id].Contents
        for e := contents.Front(); e != nil; e = e.Next() {
            oid := e.Value.(Objid)
            n := objects[oid].Name
            if compare(n, name) == 0 {
                return oid
            }
            if hasPrefix(n, name) {
                if match > 0 {
                    return Ambiguous
                }
                match = oid
            }
        }
    }
    return
}

func MatchObject(p Objid, name string) Objid {
    if len(name) == 0 {
        return Nothing
    }
    if name[0] == '#' {
        if num, err := strconv.Atoi(name[1:]); err != nil {
            return FailedMatch
        } else {
            return Objid(num)
        }
    }
    if !Valid(p) {
        return FailedMatch
    }
    if compare(name, "me") == 0 {
        return p
    }
    if compare(name, "here") == 0 {
        return ObjectLocation(p)
    }
    return matchContents(p, name)
}