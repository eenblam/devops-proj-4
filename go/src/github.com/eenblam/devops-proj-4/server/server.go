package server

import (
    "fmt"
    //"log"

    "encoding/json"
    "html/template"
    "net/http"
    "net/url"

    "github.com/go-redis/redis"
)

func RootHandler(client *redis.Client) http.HandlerFunc {
    // Return home page
    return func (w http.ResponseWriter, r *http.Request) {
        // This doesn't actually do anything at the moment, hence nil at bottom
        t, err := template.ParseFiles("views/index.html")
        if err != nil {
            fmt.Println(err)
            fmt.Fprintf(w, "Well, this is embarrassing...")
            return
        }

        t.Execute(w, nil)
    }
}

func GetJSONHandler(client *redis.Client) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        query, e := url.ParseQuery(r.URL.RawQuery)
        if e != nil {
            fmt.Println("Bad query")
        }

        // Get keys as map
        m := make(map[string]interface{})
        for key, _ := range query {
            value, e := client.Get(key).Result()
            if e != nil {
                fmt.Println("query for missing key:", key)
            }

            m[key] = value
        }

        // Serialize map to JSON
        json.NewEncoder(w).Encode(m)
    }
}

func GetHandler(client *redis.Client) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        query, e := url.ParseQuery(r.URL.RawQuery)
        if e != nil {
            fmt.Println("Bad query")
        }

        fmt.Println(r.URL.RawQuery)

        // Get keys of map query as array
        keys := make([]string, len(query))
        i := 0
        for k := range query {
            keys[i] = k
            i++
        }

        // Make all lookups at once
        mv, _ := client.MGet(keys...).Result()
        // Just dump it as a Go array
        fmt.Fprintln(w, mv)
    }
}

func SetHandler(client *redis.Client) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        query, e := url.ParseQuery(r.URL.RawQuery)
        if e != nil {
            fmt.Println("Could not parse query:", r.URL.RawQuery)
        }

	for key, values := range query {
            // Only accept the first value, I guess
            if len(values) == 0 {
                continue
            }
            value := values[0]
            err := client.Set(key, value, 0).Err()
            if err != nil {
                fmt.Println(err)
            }
        }
    }
}
