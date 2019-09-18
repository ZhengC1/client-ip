package main
import (
  "net/http"
  "strings"
)

// this code from 
// https://hackernoon.com/how-to-create-a-web-server-in-go-a064277287c9

// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}

func main() {
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
