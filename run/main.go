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
	http.HandleFunc("/", app)
	
	if err := http.ListenAndServe(":8080", nil); err != nil { //TODO: deal with $PATH variable as port number
		panic(err)
	}
}
