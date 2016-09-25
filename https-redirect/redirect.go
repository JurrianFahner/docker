package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var port string
var host string
var redirectPath bool

func createRedirectURL(r *http.Request) string {
	var redirectURL = "https://"

	if host == "" {
		host = r.Host
	}

	redirectURL = redirectURL + host

	if redirectPath {
		redirectURL = redirectURL + r.URL.Path
	}

	return redirectURL
}

func redirect(w http.ResponseWriter, r *http.Request) {
	var newURL = createRedirectURL(r)
	http.Redirect(w, r, newURL, 301)
	fmt.Println(time.Now().Format(time.RFC3339), r.Host, r.URL.Path, "->", newURL, r.UserAgent())

}

func readEnv() {
	port = os.Getenv("PORT")
	host = os.Getenv("HOST")
	redirectPath = false

	var envRedirectPath = os.Getenv("REDIRECTPATH")
	if envRedirectPath != "" {
		if strings.ToLower(envRedirectPath) == "true" {
			redirectPath = true
		}
	}

	if port == "" {
		port = "80" // default
	}
}

func main() {
	readEnv()

	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
