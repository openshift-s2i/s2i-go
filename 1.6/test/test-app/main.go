package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"runtime"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Received request from %s for %s", r.RemoteAddr, r.URL)

	h := w.Header()
	h.Set("Content-Type", "text/plain")

	fmt.Fprint(w, "Hello world!\n\n")
  fmt.Fprint(w, "oli go go openshift!\n\n")
	fmt.Fprintf(w, "Go version: %s\n", runtime.Version())
  
  //for _, e := range os.Environ() {
        //pair := strings.Split(e, "=")
    //    fmt.Fprintf(w, e)
    //}
}

func main() {
	flag.Parse()

	glog.Info("Starting server...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
