package main

import (
    "gopkg.in/mgo.v2"
    "os"
)

var session, _ = mgo.Dial(os.Getenv("MONGODB_URL"))

var Db = session.DB(os.Getenv("DB_NAME"))
var Rooms = Db.C("rooms")

type Room struct {
    RoomId string
    Text string
    Mode string
    Count int
}
