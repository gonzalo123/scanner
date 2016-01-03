package main

import (
    "log"
    "fmt"
    "net/http"
    "github.com/googollee/go-socket.io"
)

func main() {
    server, err := socketio.NewServer(nil)

    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        so.Join("messages");
    })

    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    http.Handle("/socket.io/", server)

    http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
        secret := "a60fa90b948b0f3ae204d89c01fab268"
        token := r.URL.Query().Get("token")
        if secret == r.URL.Query().Get("_secret") {
            m := make(map[string]interface{})
            m["text"] = r.URL.Query().Get("text")
            m["format"] = r.URL.Query().Get("format")
            server.BroadcastTo("messages", token, m)
            fmt.Fprintf(w, "OK")
        } else {
            fmt.Fprintf(w, "NOK")
        }
        w.Header().Set("Content-Type", "application/json")
    })

    http.Handle("/", http.FileServer(http.Dir("./asset")))

    log.Println("Serving at localhost:5000...")
    log.Fatal(http.ListenAndServe(":5000", nil))
}