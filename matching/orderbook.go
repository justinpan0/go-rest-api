package matching

import (
	logger "github.com/siddontang/go-log/log"
	//"github.com/timtadh/data-structures/hashtable"
	//"github.com/jupp0r/go-priority-queue"
)

//TODO:: Add OrderBook struct

//Order struct for order, complying 0x protocol convention
type Order struct {
	Hash                  string `json:"hash"`
	MakerAddress          string `json:"makerAddress"`
	TakerAddress          string `json:"takerAddress"`
	FeeRecipientAddress   string `json:"feeRecipientAddress"`
	SenderAddress         string `json:"senderAddress"`
	MakerAssetAmount      string `json:"makerAssetAmount"`
	TakerAssetAmount      string `json:"takerAssetAmount"`
	MakerFee              string `json:"makerFee"`
	TakerFee              string `json:"takerFee"`
	ExpirationTimeSeconds string `json:"expirationTimeSeconds"`
	Salt                  string `json:"salt"`
	MakerAssetData        string `json:"makerAssetData"`
	TakerAssetData        string `json:"takerAssetData"`
	MakerFeeAssetData     string `json:"makerFeeAssetData"`
	TakerFeeAssetData     string `json:"takerFeeAssetData"`
	ExchangeAddress       string `json:"exchangeAddress"`
	Signature             string `json:"signature"`
	TakerAssetAmountLeft  string `json:"takerAssetAmountLeft"`
}

//Orders an array of Order
type Orders []Order

var orderDB = Orders{}

//SetOrderDB add an order to orderDB
func SetOrderDB(newOrder Order) {
	orderDB = append(orderDB, newOrder)
	logger.Info("setOrderDB: adding order with hash", newOrder.Hash)
}

//GetOrderByHashDB get an order by hash
func GetOrderByHashDB(orderHash string) Order {
	for _, singleOrder := range orderDB {
		if singleOrder.Hash == orderHash {
			return singleOrder
		}
	}

	return Order{}
}

//GetOrdersDB get all orders
func GetOrdersDB() Orders {
	//TODO: Implement Paginator
	logger.Info("getOrdersDB: getting orderDB")
	return orderDB
}

//GetOrderbookDB get orderbook for specific asset pairs
func GetOrderbookDB(baseAssetData string, quoteAssetData string) (Orders, Orders) {
	//TODO: Implement Paginator
	var bids = Orders{}
	var asks = Orders{}
	for _, singleOrder := range orderDB {
		if singleOrder.MakerAssetData == baseAssetData && singleOrder.TakerAssetData == quoteAssetData {
			logger.Info("getOrdersDB: add bid", singleOrder.Hash)
			bids = append(bids, singleOrder)
		} else if singleOrder.MakerAssetData == quoteAssetData && singleOrder.TakerAssetData == baseAssetData {
			logger.Info("getOrdersDB: add ask", singleOrder.Hash)
			asks = append(asks, singleOrder)
		}
	}
	return bids, asks
}
