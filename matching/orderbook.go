package matching

import (
	//"github.com/timtadh/data-structures/hashtable"
	//"github.com/jupp0r/go-priority-queue"
)

type order struct {
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

type Orders []order
var orderDB = Orders{}

func setOrderDB(newOrder order) {	
	orderDB = append(orderDB, newOrder)
	logger.Println("setOrderDB: adding order with hash", newOrder.Hash)
}

func getOrderByHashDB(orderHash string) (order) {
	for _, singleOrder := range orderDB {
		if singleOrder.Hash == orderHash {
			return singleOrder
		}
	}

	return order{}
}

func getAssetPairsDB(assetDataA string, assetDataB string) (Orders) {
	//TODO: Implement Paginator
	var result = Orders{}
	for _, singleOrder := range orderDB {
		if singleOrder.MakerAssetData == assetDataA && singleOrder.TakerAssetData == assetDataB {
			result = append(result, singleOrder)
		}
	}
	return result
}

func getOrdersDB() (Orders) {
	//TODO: Implement Paginator
	logger.Println("getOrdersDB: getting orderDB")
	return orderDB
}

func getOrderbookDB(baseAssetData string, quoteAssetData string) (Orders, Orders) {
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