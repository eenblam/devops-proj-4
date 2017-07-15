package server

import (
    "fmt"
    //"log"

    "encoding/json"
    //"html/template"
    "net/http"
    "net/url"

    "github.com/go-redis/redis"
)

func RootHandler(client *redis.Client) http.HandlerFunc {
    // Return home page
    return func (w http.ResponseWriter, r *http.Request) {
        //data := nil

        //t, err := template.ParseFiles("index.html")
        //t.execute(w, data)
    }
}

func GetJSONHandler(client *redis.Client) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.URL.Path)

        query, e := url.ParseQuery(r.URL.RawQuery)
        if e != nil {
            fmt.Println("Bad query")
        }

        fmt.Println(r.URL.RawQuery)

        // Get keys as map
        m := make(map[string]string)
        for key, _ := range query {
            value, e := client.Get(key).Result()
            if e != nil {
                fmt.Println("I don't have", key)
            }

            m[key] = value
        }

        // Serialize map to JSON
        jsonResult, jsonErr := json.Marshal(m)
        if jsonErr != nil {
            fmt.Println("Could not convert values to JSON")
        }

        fmt.Fprintln(w, jsonResult)
    }
}

func GetHandler(client *redis.Client) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.URL.Path)

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
        fmt.Fprintln(w, mv)
    }
}

func SetHandler(client *redis.Client) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        // asdf
        fmt.Println(r.URL.Path)

        query, e := url.ParseQuery(r.URL.RawQuery)
        if e != nil {
            fmt.Println("Bad query")
        }

        fmt.Println(r.URL.RawQuery)

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
