package main

import (
    "gopkg.in/mgo.v2/bson"
    "github.com/googollee/go-socket.io"
    "log"
    "time"
    "encoding/json"
)

type socketHandlerFunc func (msg string)

func handleDisconnectionEvent (so socketio.Socket, roomId string) socketHandlerFunc {
    return func (reason string) {
        q := bson.M{"roomid": roomId}
        up := bson.M{"$inc": bson.M{"count": -1}}
        if err := Rooms.Update(q, up); err != nil {
            log.Println(err.Error())
            return
        }

        var result Room
        err := Rooms.Find(q).One(&result)
        if err != nil {
            log.Println(err.Error())
            return
        }
        if result.Count == 0 {
            // cron job for deleting the room
            time.AfterFunc(5 * time.Minute, func () {
                Rooms.Find(q).One(&result)
                if result.Count == 0 {
                    Rooms.Remove(q)
                }    
            })
        }
    }
}

func handleRoomEvent (so socketio.Socket) socketHandlerFunc {
    return func (roomId string) {
        so.Join(roomId)
        q := bson.M{"roomid": roomId}
        up := bson.M{"$inc": bson.M{"count": 1}}

        if err := Rooms.Update(q, up); err != nil {
            log.Println(err.Error())
            return 
        }

        so.On("disconnection", handleDisconnectionEvent(so, roomId))
    }
}

func handleEditEvent (so socketio.Socket) socketHandlerFunc {
    return func (change string) {
        var f interface{}
        if err := json.Unmarshal([]byte(change), &f); err != nil {
            panic(err)
        }
        m := f.(map[string]interface{})
        roomId := m["roomId"].(string)

        // save the change to the database
        text := m["text"].(string)
        q := bson.M{"roomid": roomId}
        up := bson.M{"$set": bson.M{"text": text}}
        err := Rooms.Update(q, up)
        if err != nil {
            log.Println(err.Error())
            return
        }

        // broadcast changes to necessary clients
        so.BroadcastTo(roomId, "edit", change)
    }
}

func handleSyntaxChangeEvent (so socketio.Socket) socketHandlerFunc {
    return func (change string) {
        var f interface{}
        if err := json.Unmarshal([]byte(change), &f); err != nil {
            panic(err)
        }
        m := f.(map[string]interface{})
        roomId := m["roomId"].(string)

        // save the change to the database
        mode := m["mode"].(string)
        q := bson.M{"roomid": roomId}
        up := bson.M{"$set": bson.M{"mode": mode}}
        err := Rooms.Update(q, up)
        if err != nil {
            log.Println(err.Error())
            return
        }

        so.BroadcastTo(roomId, "syntax", change)
    }
}

func SocketDef (so socketio.Socket) {
    so.On("room", handleRoomEvent(so))

    so.On("edit", handleEditEvent(so))

    so.On("syntax", handleSyntaxChangeEvent(so))
}

