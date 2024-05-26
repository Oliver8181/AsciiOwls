package main

import (
	"fmt"
	"net/http"
)

func sep(s string, r rune) (string, string) {
	for i, e := range s {
		if e == r {
			return s[:i], s[i+1:]
		}
	}
	return s, ""
}

func app(w http.ResponseWriter, r *http.Request) {
	_, query := sep(r.URL.Path, '?')
	fmt.Fprintf(w, "<DOCTYPE html><html>"+query+"</html>")

}

func main() {
        log.Print("starting server...")
        http.HandleFunc("/", app)

        // Determine port for HTTP service.
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("defaulting to port %s", port)
        }

        // Start HTTP server.
        log.Printf("listening on port %s", port)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal(err)
        }
}
