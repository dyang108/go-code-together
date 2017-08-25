package main

import (
    "gopkg.in/mgo.v2"
    "os"
)

var session, _ = mgo.Dial(os.Getenv("MONGODB_URI"))

var Db = session.DB("")
var Rooms = Db.C("rooms")

type Room struct {
    RoomId string
    Text string
    Mode string
    Count int
}
