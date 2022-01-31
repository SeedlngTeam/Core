package main

import (
	"fmt"
	"net/http"
)

func catchAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to Seedlng!!")
}

func main() {
	http.HandleFunc("/", catchAll)
	fmt.Print("Server Staring on :3000 ...")
	http.ListenAndServe("https://localhost:3000", nil)
}
