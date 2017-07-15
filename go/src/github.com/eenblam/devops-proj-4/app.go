package main

import (
    "log"

    "net/http"

    "github.com/eenblam/devops-proj-4/redisclient"
    "github.com/eenblam/devops-proj-4/server"
)

func main() {
    client := redisClient.NewClient()
    defer client.Close()

    http.HandleFunc("/", server.RootHandler(client))
    http.HandleFunc("/api/get", server.GetHandler(client))
    http.HandleFunc("/api/getJSON", server.GetJSONHandler(client))
    http.HandleFunc("/api/set", server.SetHandler(client))
    log.Fatal(http.ListenAndServe(":5000", nil))
}
