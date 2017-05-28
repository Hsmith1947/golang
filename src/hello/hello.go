package main

import (
    "html/template"
    "net/http"
    "hello/stringutil"
)

var templates map[string]*template.Template
func init() {
    templates = make(map[string]*template.Template)
    templates["hello"] = template.Must(template.ParseFiles("../src/hello/templates/base.html", "../src/hello/templates/hello.html"))
    templates["test"] = template.Must(template.ParseFiles("../src/hello/templates/base.html", "../src/hello/templates/test.html"))
}

func main() {
    http.HandleFunc("/", handler)
    
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    name := stringutil.AddPeriod(r.URL.Path[1:])

    if name == "test" {
        templates["test"].Execute(w, map[string]interface{}{})
        return
    }

    templates["hello"].Execute(w, map[string]interface{}{
        "Name": name,
    })
}