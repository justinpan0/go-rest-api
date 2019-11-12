package main

import (
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/siddontang/go-log/log"
	"github.com/zimengpan/go-rest-api/matching"
)

func main() {
	initValidator()
	matching.StartEngine()

	logger.Info("Boomflow starts")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/v1/order/{productId}", setOrder)
	router.HandleFunc("/v1/orders/{orderHash}", getOrderByHash)
	router.HandleFunc("/v1/orders", getOrders)
	router.HandleFunc("/v1/orderbook", getOrderbook)

	logger.Fatal(http.ListenAndServe(":9000", router))
}
