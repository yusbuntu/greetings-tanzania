package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
)

func main() {
    // Read environment variables
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    greeting := os.Getenv("GREETING")
    if greeting == "" {
        greeting = "Greetings from Kilimanjaro"
    }

    // Serve static files
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Handle root route
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("static/index.html"))
        data := map[string]string{
            "Greeting": greeting,
        }
        tmpl.Execute(w, data)
    })

    fmt.Printf("Listening on port %s...\n", port)
    http.ListenAndServe(":"+port, nil)
}

