// MIT Licence

// Info on package

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var (
	sourceAddress, targetAddress string
)

func main() {

	targetAddress = "http://127.0.0.1:4000"
	sourceAddress = ":8000"

	slowPaths := map[string]bool{
		"/wp-admin/":                    true,
		"/wp-login.php":                 true,
		"/phpMyAdmin/scripts/setup.php": true,
	}

	target, err := url.Parse(targetAddress)
	if err != nil {
		log.Fatal("Bad destination.")
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, proxy, slowPaths)
	})
	http.ListenAndServe(sourceAddress, nil)
}

func handler(w http.ResponseWriter, r *http.Request, proxy *httputil.ReverseProxy, slowPaths map[string]bool) {
	if slowPaths[r.URL.Path] {
		time.Sleep(5000 * time.Millisecond)

		w.WriteHeader(573)
		w.Write([]byte("Error.\n"))

	} else {
		proxy.ServeHTTP(w, r)
	}

}
