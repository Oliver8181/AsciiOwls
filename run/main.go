package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func sep(s string, r rune) (string, string) {
	for i, e := range s {
		if e == r {
			return s[:i], s[i+1:]
		}
	}
	return s, ""
}

func generate() string {
	names := []string{"oliver", "mathew", "bob"}
	rand := rand.New(rand.NewSource(time.Now().UnixNano() + 32383))
	return names[rand.Int31n(int32(len(names)))]
}

func wrap(head, body string) string {
	files, _ := filepath.Glob("*")
	return "<!DOCTYPE html><html><head>" + head + "</head><body>" + body + "<br>" + strings.Join(files, ", ") + "</body></html>"
}

func app(w http.ResponseWriter, r *http.Request) {
	path, _ := sep(r.URL.Path, '?')
	data, err := os.ReadFile("site/" + path)
	if err != nil {
		name, err := r.Cookie("__session")
		if err != nil {
			name := generate()
			http.SetCookie(w, &http.Cookie{
				Name:  "__session",
				Value: name,

				// Expires:  time.Time,

				// MaxAge=0 means no 'Max-Age' attribute specified.
				// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
				// MaxAge>0 means Max-Age attribute present and given in seconds
				MaxAge:   24 * 60 * 60,
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteStrictMode,
			})

			fmt.Fprintf(w, wrap("", "We're going to call you "+name))
		} else {
			fmt.Fprintf(w, wrap("", "Welcome back "+name.Value))
		}
	} else {
		fmt.Fprintf(w, "%s", data)
	}

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	return port
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", app)

	// Determine port for HTTP service.
	port := getPort() // test on http://127.0.0.1:8080/
	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
