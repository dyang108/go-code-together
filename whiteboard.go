package main

import (
    "html/template"
    "net/http"
    "os"
    "strings"
    "gopkg.in/mgo.v2/bson"
)

func WhiteboardIndex (w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        return
    }
    tmplVars := IndexTmpl{
        Socket: os.Getenv("SOCKET_URL"),
        Env: os.Getenv("NODE_ENV"),
        BaseUrl: os.Getenv("BASE_URL"),
    }
    t, _ := template.ParseFiles("index-whiteboard.html")
    t.Execute(w, tmplVars)
}

func CreateWhiteboard (roomChannels *map[string]chan bson.M) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        // check if correct method
        if r.Method != "POST" {
            return
        }

        // need to parse the form in order to get data
        r.ParseForm()
        roomId := "whiteboard/" + strings.Join(r.Form["roomId"], "")
        if strings.Contains(roomId, " ") {
            http.Redirect(w, r, os.Getenv("BASE_URL"), 301)
            return
        }


        var result Room
        if err := Rooms.Find(bson.M{"roomid": roomId}).One(&result); err != nil {
            if err.Error() == "not found" {
                newSt := Room{
                    RoomId: roomId,
                    Text: "// type code here",
                    Mode: "ace/mode/javascript",
                    Count: 0,
                }
                if err := Rooms.Insert(&newSt); err != nil {
                    http.Error(w, "Error occurred when inserting in database " + err.Error(), 501)
                    return
                }
                // initialize the channel for the room
                (*roomChannels)[roomId] = make(chan bson.M)
                go DigestEvents((*roomChannels)[roomId], roomId)
            } else {
                http.Error(w, "Error occurred when querying database " + err.Error(), 501)
                return
            }
        }
        http.Redirect(w, r, os.Getenv("BASE_URL") + roomId, 301)
    }
}
