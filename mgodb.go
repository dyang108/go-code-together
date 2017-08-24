package main

import (
    "gopkg.in/mgo.v2"
)

var session, _ = mgo.Dial("mongodb://localhost:27017/")

var Db = session.DB("webcoder")
var Rooms = Db.C("rooms")

type Room struct {
    RoomId string
    Text string
    Mode string
    Count int
}
