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

    // Liveness probe
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })

    // Readiness probe
    http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
        // Here you can add checks to verify if your app is ready
        // For simplicity, I'm just returning OK for now
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ready"))
    })

    fmt.Printf("Listening on port %s...\n", port)
    http.ListenAndServe(":"+port, nil)
}

