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

        so.BroadcastTo(roomId, "otherUserDisconnect", so.Id())

        var result Room
        if err := Rooms.Find(q).One(&result); err != nil {
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
        go so.BroadcastTo(roomId, "countChange", "1")
        go so.Emit("countChange", "1")
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
        // broadcast changes to necessary clients
        go so.BroadcastTo(roomId, "edit", change)

        // save the change to the database
        text := m["text"].(string)
        q := bson.M{"roomid": roomId}
        up := bson.M{"$set": bson.M{"text": text}}
        if err := Rooms.Update(q, up); err != nil {
            log.Println(err.Error())
            return
        }
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
        go so.BroadcastTo(roomId, "syntaxChange", change)

        // save the change to the database
        mode := m["mode"].(string)
        q := bson.M{"roomid": roomId}
        up := bson.M{"$set": bson.M{"mode": mode}}
        if err := Rooms.Update(q, up); err != nil {
            log.Println(err.Error())
            return
        }
    }
}

func handleCursorChangeEvent (so socketio.Socket) socketHandlerFunc {
    return func (change string) {
        so.BroadcastTo(so.Rooms()[0], "changeCursor", change)
    }
}

func handleSelectionChangeEvent (so socketio.Socket) socketHandlerFunc {
    return func (change string) {
        so.BroadcastTo(so.Rooms()[0], "changeSelection", change)
    }
}

type socketHandlerFuncWithAck func (msg string) int

func handleRoomNameEditEvent (so socketio.Socket) socketHandlerFuncWithAck {
    return func (roomId string) int {
        var result Room
        q := bson.M{"roomid": roomId}
        if err := Rooms.Find(q).One(&result); err != nil {
            return 0
        } else {
            return result.Count
        }
    }
}

func SocketDef (so socketio.Socket) {
    so.On("room", handleRoomEvent(so))
    so.On("edit", handleEditEvent(so))
    so.On("syntaxChange", handleSyntaxChangeEvent(so))
    so.On("changeSelection", handleSelectionChangeEvent(so))
    so.On("changeCursor", handleCursorChangeEvent(so))
    so.On("roomNameEdit", handleRoomNameEditEvent(so))
}
