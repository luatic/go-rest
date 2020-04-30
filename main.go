package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	http.ListenAndServe(":1212", router)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is a test !!!"))
}
