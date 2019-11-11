package main

import (
	"os"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var logger = log.New(os.Stderr, "Main\t", log.Lshortfile)

func main() {
	initValidator()

	logger.Println("Boomflow starts")

	router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/v1/asset_pairs")
	//router.HandleFunc("/v1/orders")
	
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/v1/order", setOrder)
	router.HandleFunc("/v1/order/{orderHash}", getOrderByHash)
	router.HandleFunc("/v1/orders", getOrders)
	router.HandleFunc("/v1/orderbook", getOrderbook)

	log.Fatal(http.ListenAndServe(":9000", router))
}