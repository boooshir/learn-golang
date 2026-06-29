package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var addr *string = flag.String("addr", ":3000", "http service address")

func main() {
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)
	http.ListenAndServe(*addr, nil)
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	tokens := []string{"this", "is", "a", "live", "event", "test", "from", "dgd", "warsaw", "2024"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		// fmt.Fprintf(w, "%s\n", content)
		w.(http.Flusher).Flush()

		// intentional wait
		time.Sleep(time.Millisecond * 420)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
