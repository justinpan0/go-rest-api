package matching

import (
	"os"
	"log"
	//"github.com/timtadh/data-structures/hashtable"
	//"github.com/jupp0r/go-priority-queue"
)

var logger = log.New(os.Stderr, "Main\t", log.Lshortfile)

type Order struct {
	Hash					string	`json:"hash"`
	MakerAddress			string 	`json:"makerAddress"`
	TakerAddress 			string 	`json:"takerAddress"`
	FeeRecipientAddress		string 	`json:"feeRecipientAddress"`
	SenderAddress 			string 	`json:"senderAddress"`
	MakerAssetAmount		string 	`json:"makerAssetAmount"`
	TakerAssetAmount		string 	`json:"takerAssetAmount"`
	MakerFee				string 	`json:"makerFee"`
	TakerFee				string 	`json:"takerFee"`
	ExpirationTimeSeconds	string 	`json:"expirationTimeSeconds"`
	Salt					string 	`json:"salt"`
	MakerAssetData			string 	`json:"makerAssetData"`
	TakerAssetData			string 	`json:"takerAssetData"`
	MakerFeeAssetData		string 	`json:"makerFeeAssetData"`
	TakerFeeAssetData		string 	`json:"takerFeeAssetData"`
	ExchangeAddress			string 	`json:"exchangeAddress"`
	Signature 				string 	`json:"signature"`
	TakerAssetAmountLeft	string	`json:"takerAssetAmountLeft"`
}

type Orders []Order
var orderDB = Orders{}

func SetOrderDB(newOrder Order) {	
	orderDB = append(orderDB, newOrder)
	logger.Println("setOrderDB: adding order with hash", newOrder.Hash)
}

func GetOrderByHashDB(orderHash string) (Order) {
	for _, singleOrder := range orderDB {
		if singleOrder.Hash == orderHash {
			return singleOrder
		}
	}

	return Order{}
}

func GetAssetPairsDB(assetDataA string, assetDataB string) (Orders) {
	//TODO: Implement Paginator
	var result = Orders{}
	for _, singleOrder := range orderDB {
		if singleOrder.MakerAssetData == assetDataA && singleOrder.TakerAssetData == assetDataB {
			result = append(result, singleOrder)
		}
	}
	return result
}

func GetOrdersDB() (Orders) {
	//TODO: Implement Paginator
	logger.Println("getOrdersDB: getting orderDB")
	return orderDB
}

func GetOrderbookDB(baseAssetData string, quoteAssetData string) (Orders, Orders) {
	//TODO: Implement Paginator
	var bids = Orders{}
	var asks = Orders{}
	for _, singleOrder := range orderDB {
		if singleOrder.MakerAssetData == baseAssetData && singleOrder.TakerAssetData == quoteAssetData {
			logger.Println("getOrdersDB: add bid", singleOrder.Hash)
			bids = append(bids, singleOrder)
		} else if singleOrder.MakerAssetData == quoteAssetData && singleOrder.TakerAssetData == baseAssetData {
			logger.Println("getOrdersDB: add ask", singleOrder.Hash)
			asks = append(asks, singleOrder)
		}
	}
	return bids, asks
}