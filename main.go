package main

import (
    "fmt"
    "net/http"
    "time"
)

func hello(w http.ResponseWriter, req *http.Request) {
    time.Sleep(1 * time.Second)
    w.Write([]byte(`
                <!DOCTYPE html>
                <html>
                <head>
                <title>Hello WG Forge</title>
                <script src="statics/jquery-1.12.4.min.js"></script>
                <link href="statics/style.css" rel="stylesheet">
                </head>
                <body>
                <h1>Hello WG Forge 2020</h1>
                </body>
                </html>`))
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8080", nil)
}
