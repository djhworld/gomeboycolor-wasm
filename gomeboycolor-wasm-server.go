// +build !wasm
//go:generate statik -src=static/gomeboycolor

package main

import (
	"fmt"
	_ "github.com/djhworld/gomeboycolor-wasm/statik"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	fmt.Println("Open your web browser and navigate to: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
