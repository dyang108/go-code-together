package main

import (
    "html/template"
    "net/http"
    "log"
    "strings"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "os"
)

type EditorTmpl struct {
    Mode string
    Socket string
    Text string
    Env string
    Count int
}

type IndexTmpl struct {
    BaseUrl string
    Env string
    Socket string
}

func displayEditor(w http.ResponseWriter, r *http.Request, path string) {
    var result Room
    // TODO: block until channel is empty
    if err := Rooms.Find(bson.M{"roomid": path}).One(&result); err != nil {
        if err.Error() == "not found" {
            http.Redirect(w, r, os.Getenv("BASE_URL"), 301)
        } else {
            http.Error(w, "Error occurred when querying database " + err.Error(), 501)
        }
        return
    } else {
        tmplVars := EditorTmpl{
            Mode: result.Mode,
            Socket: os.Getenv("SOCKET_URL"),
            Text: result.Text,
            Env: os.Getenv("NODE_ENV"),
            Count: result.Count,
        }
        w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
        t, _ := template.ParseFiles("editor.html")
        t.Execute(w, tmplVars)
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    if path := r.URL.Path[1:]; len(path) != 0 {
        displayEditor(w, r, path)
        return
    }
    tmplVars := IndexTmpl{
        Socket: os.Getenv("SOCKET_URL"),
        Env: os.Getenv("NODE_ENV"),
        BaseUrl: os.Getenv("BASE_URL"),
    }
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, tmplVars)
}

func createRoom (roomChannels *map[string]chan bson.M) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        // check if correct method
        if r.Method != "POST" {
            return
        }

        // need to parse the form in order to get data
        r.ParseForm()
        roomId := strings.Join(r.Form["roomId"], "")
        if strings.Contains(roomId, " ") || roomId == "about" {
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

func main() {
    roomChannels := make(map[string]chan bson.M)

    io := InitSocket(&roomChannels)
    http.HandleFunc("/socket.io/", func (w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        io.ServeHTTP(w, r)
    })

    Rooms.EnsureIndex(mgo.Index{
        Key: []string{"roomid"},
        Unique: true,
        DropDups: true,
        Background: true,
    })

    // references in templates load to /assets
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    http.HandleFunc("/", index)
    http.HandleFunc("/create-room", createRoom(&roomChannels))

    log.Println("Starting server on port " + os.Getenv("PORT"))
    http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
