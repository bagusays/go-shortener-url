package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	db := DbConn()

	Router(r, db)

	err := http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatal()
	}

	log.Println("listen and serve on 0.0.0.0:8888 (for windows localhost:8888)")
}
