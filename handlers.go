package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/zimengpan/go-rest-api/matching"

	"github.com/gorilla/mux"
)

func setOrder(w http.ResponseWriter, r *http.Request) {
	//TODO: http request error code & handling
	var newOrder matching.Order
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Println("setOrder: error reading data")
		return;
	} else {
		// validate order
		if err, _ := rs.ValidateBytes(reqBody); len(err) > 0 {
			logger.Println("setOrder: invalid order data")
			return
		}
	}

	json.Unmarshal(reqBody, &newOrder)
	logger.Println("setOrder: creating order with hash", newOrder.Hash)
	matching.SetOrderDB(newOrder)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

func getOrderByHash(w http.ResponseWriter, r *http.Request) {
	//TODO: http request error code & handling
	orderHash := mux.Vars(r)["orderHash"]
	logger.Println("getOrderByHash: get order", orderHash)
	
	result := matching.GetOrderByHashDB(orderHash)
	json.NewEncoder(w).Encode(result)
}

func getAssetPairs(w http.ResponseWriter, r *http.Request) {
	//TODO: http request error code & handling
	assetDataA := r.URL.Query().Get("assetDataA")
	assetDataB := r.URL.Query().Get("assetDataB")
	logger.Println("getOrderbook: get the orderbook for\n\tassetDataA:", assetDataA)

	result := matching.GetAssetPairsDB(assetDataA, assetDataB)
	json.NewEncoder(w).Encode(result)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	//TODO: http request error code & handling
	logger.Println("getOrderByHash: get the all orders within criteria")
	
	result := matching.GetOrdersDB()
	json.NewEncoder(w).Encode(result)
}

func getOrderbook(w http.ResponseWriter, r *http.Request) {
	//TODO: http request error code & handling
	baseAssetData := r.URL.Query().Get("baseAssetData")
	quoteAssetData := r.URL.Query().Get("quoteAssetData")
	logger.Println("getOrderbook: get the orderbook for\n\tbaseAssetData:", baseAssetData)

	bids, asks := matching.GetOrderbookDB(baseAssetData, quoteAssetData)
	result := map[string]matching.Orders{"bids": bids, "asks": asks}
	json.NewEncoder(w).Encode(result)
}

func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome home!")
}