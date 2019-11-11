package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"github.com/zimengpan/go-rest-api/matching"
	"github.com/zimengpan/go-rest-api/service"
)

var productId2Writer sync.Map

func getWriter(productId string) *kafka.Writer {
	writer, found := productId2Writer.Load(productId)
	if found {
		return writer.(*kafka.Writer)
	}

	newWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{"localhost:9092"},
		Topic:        "matching_order_" + productId,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 5 * time.Millisecond,
	})
	productId2Writer.Store(productId, newWriter)
	return newWriter
}

func setOrder(w http.ResponseWriter, r *http.Request) {
	//TODO: http request error code & handling
	productId := mux.Vars(r)["productId"]
	var newOrder matching.Order
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Fatalln("setOrder: error reading data")
		return
	} else {
		// validate order
		if err, _ := rs.ValidateBytes(reqBody); len(err) > 0 {
			logger.Fatalln("setOrder: invalid order data")
			return
		}
	}

	json.Unmarshal(reqBody, &newOrder)
	logger.Println("setOrder: submit order with hash", newOrder.Hash)
	product, err := service.GetProductById(productId)
	if (newOrder.MakerAssetData != product.BaseAssetData || newOrder.TakerAssetData != product.QuoteAssetData) && (newOrder.TakerAssetData != product.BaseAssetData || newOrder.MakerAssetData != product.QuoteAssetData) {
		logger.Fatalln("setOrder: productId and asset pairs unmatched")
		return
	}
	logger.Println("setOrder: pair ", product.BaseCurrency, product.QuoteCurrency)

	err = getWriter(productId).WriteMessages(context.Background(), kafka.Message{Value: reqBody})
	if err != nil {
		logger.Fatalln(err)
	}

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

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
