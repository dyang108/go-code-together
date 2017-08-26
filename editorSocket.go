package main

import (
    "gopkg.in/mgo.v2/bson"
    "github.com/googollee/go-socket.io"
    "log"
    "time"
    "encoding/json"
)

type socketHandlerFunc func (msg string)
type socketHandlerFuncWithAck func (msg string) int

func handleDisconnectionEvent (so socketio.Socket, roomChannel chan bson.M, roomId string) socketHandlerFunc {
    return func (reason string) {
        so.BroadcastTo(roomId, "otherUserDisconnect", so.Id())

        go func () {
            // send message to update db
            roomChannel <- bson.M{"$inc": bson.M{"count": -1}}
            q := bson.M{"roomid": roomId}
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
        }()
    }
}

func handleEditEvent (so socketio.Socket, roomChannel chan bson.M, roomId string) socketHandlerFunc {
    return func (change string) {
        // broadcast changes immediately to necessary clients
        so.BroadcastTo(roomId, "edit", change)
        var f interface{}
        if err := json.Unmarshal([]byte(change), &f); err != nil {
            panic(err)
        }
        m := f.(map[string]interface{})

        // launch this because we dont want to block the main goroutine with the assignment to the channel
        go func () {
            // save the change to the database
            text := m["text"].(string)
            roomChannel <- bson.M{"$set": bson.M{"text": text}}
        }()
    }
}

func handleSyntaxChangeEvent (so socketio.Socket, roomChannel chan bson.M, roomId string) socketHandlerFunc {
    return func (change string) {
        so.BroadcastTo(roomId, "syntaxChange", change)
        var f interface{}
        if err := json.Unmarshal([]byte(change), &f); err != nil {
            panic(err)
        }
        m := f.(map[string]interface{})

        // launch this because we dont want to block the main goroutine with assignment to channel
        go func () {
            // save the change to the database
            mode := m["mode"].(string)
            roomChannel <- bson.M{"$set": bson.M{"mode": mode}}
        }()
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

func InitSocket (roomChannels *map[string]chan bson.M) *socketio.Server {
    io, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    io.On("connection", func (so socketio.Socket) {
        so.On("room", func (roomId string) {
            so.BroadcastTo(roomId, "countChange", "1")
            so.Emit("countChange", "1")
            so.Join(roomId)
            if _, ok := (*roomChannels)[roomId]; !ok {
                (*roomChannels)[roomId] = make (chan bson.M)
                go DigestEvents((*roomChannels)[roomId], roomId)
            }
            roomChannel := (*roomChannels)[roomId]
            go func () {
                roomChannel <- bson.M{"$inc": bson.M{"count": 1}}
            }()
            so.On("disconnection", handleDisconnectionEvent(so, roomChannel, roomId))
            so.On("edit", handleEditEvent(so,roomChannel, roomId))
            so.On("syntaxChange", handleSyntaxChangeEvent(so, roomChannel, roomId))
        })
        so.On("changeSelection", handleSelectionChangeEvent(so))
        so.On("changeCursor", handleCursorChangeEvent(so))
        so.On("roomNameEdit", handleRoomNameEditEvent(so))
    })
    return io
}

// handles changes to this specific room, one change at a time
func DigestEvents (roomChannel chan bson.M, roomId string) {
    q := bson.M{"roomid": roomId}
    // isBusy := false
    // TODO: ensure we always make the change if it's the last one in the channel (intermediates arent as important)
    for {
        select {
        case up := <- roomChannel:
            go func () {
                // isBusy = true
                if err := Rooms.Update(q, up); err != nil {
                    log.Println(err.Error())
                }
                // isBusy = false
            }()
        }
    }
}
