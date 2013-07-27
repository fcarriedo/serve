package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.Int("p", 8080, "the port on where to serve files")
var dir = flag.String("dir", ".", "The directory to serve files from")

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Serving files from '%s' on port %d\n", *dir, *port)
	panic(http.ListenAndServe(addr, http.FileServer(http.Dir(*dir))))
}
