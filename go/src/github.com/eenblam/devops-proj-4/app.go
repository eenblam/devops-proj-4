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

    /*
    e := client.Set("key", "value", 0).Err()
    if e != nil {
        log.Fatal(e)
    }

    v, e := client.Get("key").Result()
    if e != nil {
        log.Fatal(e)
    }
    fmt.Println("key", v)
    */

    http.HandleFunc("/", server.RootHandler(client))
    http.HandleFunc("/api/get", server.GetHandler(client))
    http.HandleFunc("/api/getJSON", server.GetJSONHandler(client))
    http.HandleFunc("/api/set", server.SetHandler(client))
    log.Fatal(http.ListenAndServe(":5000", nil))
}
