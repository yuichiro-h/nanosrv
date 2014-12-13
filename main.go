package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var root string
	var port int

	flag.StringVar(&root, "root", "./", "document root.")
	flag.IntVar(&port, "port", 9876, "port")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(root)))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
