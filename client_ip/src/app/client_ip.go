package main
import (
  "net/http"
  "log"
)

// the skeleton code is from
// https://hackernoon.com/how-to-create-a-web-server-in-go-a064277287c9
// with some guidance from
// https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

func getIp(w http.ResponseWriter, r *http.Request) {
  // tested the first if on the command line
  // curl --header "X-Forwarded-For: 192.168.0.2" http://localhost:8080
  message := "Current Ip: "
  ipAddress := r.Header.Get("X-Real-Ip")
  if ipAddress == "" {
      ipAddress = r.Header.Get("X-Forwarded-For")
  }
  if ipAddress == "" {
      ipAddress = r.RemoteAddr
  }
  w.Write([]byte(message + ipAddress))
}

func main() {
  log.Println("Starting Server")
  http.HandleFunc("/", getIp)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
