package main

import (
	"net/http"

	"github.com/jamesstocktonj1/componentize-sdk/p3/net/wasihttp"
)

func init() {
	wasihttp.HandleFunc(handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

}

func main() {}
