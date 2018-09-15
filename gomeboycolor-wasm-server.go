// +build !wasm
//go:generate statik -src=static/gomeboycolor

package main

import (
	"fmt"
	"log"
    "net/http"
	"github.com/rakyll/statik/fs"
	_ "github.com/djhworld/gomeboycolor-wasm/statik"
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
