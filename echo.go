// httpserver.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Request struct {
	Host   string
	Header map[string][]string
}

func reqToString(r http.Request) string {
	req := struct {
		Host          string
		Method        string
		URL           string
		Proto         string
		RemoteAddress string
		FormValues    map[string][]string
		PostValues    map[string][]string
		Header        map[string][]string
	}{r.Host, r.Method, r.URL.String(), r.Proto, r.RemoteAddr, r.Form, r.PostForm, r.Header}

	b, _ := json.Marshal(req)

	return string(b)
}

func main() {
	addr := flag.String("addr", "localhost:8080", "where to listen for connection (default: localhost:8080 or localhost:8443 when tls is enabled)")
	tls := flag.Bool("tls", false, "listen via tls")
	flag.Parse()

	fmt.Printf("listening for requests on: %s / ssl mode:%t \n", *addr, *tls)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "%s\n", reqToString(*r))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, reqToString(*r))
	})

	if *tls {
		http.ListenAndServeTLS(*addr, "cert.pem", "key.pem", nil)
	} else {
		log.Fatal(http.ListenAndServe(*addr, nil))
	}
}
