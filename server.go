package main

import (
    "html/template"
    "net/http"
    "log"
    "strings"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/googollee/go-socket.io"
)

type EditorTmpl struct {
    Mode string
    Socket string
    Text string
    Env string
}

func displayEditor(w http.ResponseWriter, r *http.Request, path string) {
    var result Room
    err := Rooms.Find(bson.M{"roomid": path}).One(&result)
    if err != nil {
        if err.Error() == "not found" {
            http.Redirect(w, r, "http://localhost:8000/", 301)
        } else {
            http.Error(w, "Error occurred when querying database", 501)
        }
        return
    } else {
        tmplVars := EditorTmpl{
            Mode: result.Mode,
            Socket: "ws://localhost:8000/",
            Text: result.Text,
            Env: "dev",
        }
        t, _ := template.ParseFiles("editor.html")
        t.Execute(w, tmplVars)
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    if path := r.URL.Path[1:]; len(path) != 0 {
        displayEditor(w, r, path)
        return
    }
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, "http://localhost:8000/")
}

func createRoom(w http.ResponseWriter, r *http.Request) {
    // check if correct method
    if r.Method != "POST" {
        http.Error(w, "Invalid request method.", 405)
        return
    }

    // need to parse the form in order to get data
    r.ParseForm()
    roomId := strings.Join(r.Form["roomId"], "")
    if strings.Contains(roomId, " ") {
        http.Redirect(w, r, "http://localhost:8000/", 301)
        return
    }

    var result Room
    err := Rooms.Find(bson.M{"roomid": roomId}).One(&result)
    if err != nil {
        if err.Error() == "not found" {
            newSt := Room{
                RoomId: roomId,
                Text: "// type code here",
                Mode: "ace/mode/javascript",
                Count: 0,
            }
            err := Rooms.Insert(&newSt)
            if err != nil {
                http.Error(w, "Error occurred when inserting in database", 501)
            }
         } else {
            http.Error(w, "Error occurred when querying database", 501)
        }
    }
    http.Redirect(w, r, "http://localhost:8000/" + roomId, 301)
}

func main() {
    io, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }

    io.On("connection", SocketDef)
    http.Handle("/socket.io/", io)

    Rooms.EnsureIndex(mgo.Index{
        Key: []string{"roomid"},
        Unique: true,
        DropDups: true,
        Background: true,
    })

    // references in templates load to /assets
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    http.HandleFunc("/", index)
    http.HandleFunc("/create-room", createRoom)
    log.Println("Starting server on port 8000")
    http.ListenAndServe(":8000", nil)
}
