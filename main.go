package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {

	targetUrl, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)

	handler := func(w http.ResponseWriter, r *http.Request) {

		hp := strings.Split(r.Host, ":")
		h := hp[0]
		p := hp[1]

		log.Println("main:", h, p)

		if strings.Compare(h, "app.node.pi.g") == 0 {
			r.Host = "app.node.pi.g"
			log.Println("r : ", r.Host)
			proxy.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusBadGateway)
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8090", nil))
}
