package main

import (
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO FROM GO HTTP SERVER"))
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	msg := os.Getenv("MSG")
	w.Write([]byte(msg))
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Println("server started on port :9000")

	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}
