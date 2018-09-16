// +build !wasm
//go:generate statik -src=static/gomeboycolor

package main

import (
	"flag"
	"fmt"
	_ "github.com/djhworld/gomeboycolor-wasm/statik"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

var port *int = flag.Int("p", 8080, "the port to run the server on")

func main() {
	flag.Parse()

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	fmt.Printf("Open your web browser and navigate to: http://localhost:%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)

}
