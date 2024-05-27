package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var replace = map[string]string{
	"":     ".html",
	".htm": ".html",
}

func pathFilter(path string) string {
	extension := filepath.Ext(path)
	if replacement, ok := replace[extension]; ok {
		path = path[:len(path)-len(extension)] + replacement
	}
	return path
}

const pic = `
	.\//_
`

func app(w http.ResponseWriter, r *http.Request) {
	// if !filepath.IsAbs(r.URL.Path) {
	// 	fmt.Println("error")
	// 	return
	// }
	path := pathFilter(r.URL.Path)
	// insert.html if there is an error
	data, err := os.ReadFile("site/" + path)
	if err != nil {
		data, _ = os.ReadFile("site/404.html")
	}
	fmt.Fprintf(w, "%s", data)
	if filepath.Ext(path) == ".html" {
		curr := time.Now()
		fmt.Fprintf(w, "<!--%sos: %s,\nport: %s,\ntime: %s\n-->", pic, runtime.GOOS, os.Getenv("PORT"), curr.Format("2006-Jan-02 at 15:04:05 in timezone MST"))
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
