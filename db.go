package main

import (
    "sync"
)

var objects = map[Objid]*object { }

type object struct {
    id Objid
    owner Objid
    location Objid
    contents []Objid
    parent Objid
    name string
    sync.Mutex
}

var mu sync.Mutex

func newObject() Objid {
    mu.Lock()
    defer mu.Unlock()
    id := Objid(len(objects))
    objects[id] = &object{
        id: id,
        name: "",
        location: Nothing,
        parent: Nothing,
    }
    return id
}

func valid(id Objid) bool {
    _, ok := objects[id]
    return ok
}