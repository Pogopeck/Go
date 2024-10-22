package main

import (
    "html/template"
    "log"
    "net/http"
)

type Option struct {
    Value string
    Text  string
}

func main() {
    http.HandleFunc("/", homeHandler)
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    options := []Option{
        {Value: "1", Text: "Option 1"},
        {Value: "2", Text: "Option 2"},
        {Value: "3", Text: "Option 3"},
    }

    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, options)
}