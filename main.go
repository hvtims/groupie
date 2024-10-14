package main

import (
	"fmt"
	"net/http"
	"os"

	"mrg/mrg"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("check args")
		return
	}
	http.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("frontend"))))
	fmt.Println("this is your port : http://localhost:8080/ ")
	http.HandleFunc("/", mrg.HandleHome)
	http.HandleFunc("/Artist/{id}", mrg.HandlePage)
	http.ListenAndServe(":8080", nil)
}
