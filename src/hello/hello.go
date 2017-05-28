package main

import (
    "html/template"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Path[1:]
    t, _ := template.ParseFiles("templates/base.html", "templates/hello.html")
    t.Execute(w, map[string]interface{}{
        "Name": name,
    })
}