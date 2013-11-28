package main

import (
    "container/list"
)

type Object struct {
    Id Objid
    Name string
    Location Objid
    Contents *list.List
}

var objects = make(map[Objid]*Object)

func findObject(id Objid) *Object {
    obj, ok := objects[id]
    if ok {
        return obj
    }
    return nil
}

func Valid(id Objid) bool {
    return findObject(id) != nil
}

func CreateObject() Objid {
    n := len(objects)
    id := Objid(n)
    objects[id] = &Object{id, "", Nothing, list.New()}
    return id
}

func ObjectLocation(id Objid) Objid {
    return objects[id].Location
}

func ChangeLocation(id Objid, location Objid) {
    var contents *list.List
    oldLocation := objects[id].Location
    if Valid(oldLocation) {
        var e *list.Element
        contents = objects[oldLocation].Contents
        for e = contents.Front(); e != nil; e = e.Next() {
            if e.Value == id {
                break
            }
        }
        if e != nil {
            contents.Remove(e)
        }
    }
    if Valid(location) {
        objects[location].Contents.PushBack(id)
    }
    objects[id].Location = location   
}