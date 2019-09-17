package main

import (
    "fmt"
    "net/http"
)


func main() {
}

func ProcessDynamicRequests() {
   http.HandleFunc("/",
      func(w http.ResponseWriter, r *http.Request) {
          fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
      }
  )
  return true
}

func ServeStaticAssets() {
    // fs hold the instance of the static assets internal server
    fs := http.FileServer(http.Dir("static/"))
    return fs
}

func ProcessStaticRequests() {
    http.HandleFunc(
        "/static/",
        http.StripPrefix("/static/",
        server_static_assets())
    )
}

func AcceptConnections() {
    http.ListenAndServe(":3001", nil)
}
