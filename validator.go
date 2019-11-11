package main

import (
	"encoding/json"
	"github.com/qri-io/jsonschema"
)

var schemaData = []byte(`{
	"title": "Order",
	"type": "object",
	"properties": {
		"hash": {
			"type": "string"
		},
		"makerAddress": {
			"type": "string"
		},
		"takerAddress": {
			"type": "string"
		},
		"feeRecipientAddress": {
			"type": "string"
		},
		"senderAddress": {
			"type": "string"
		},
		"makerAssetAmount": {
			"type": "string"
		},
		"takerAssetAmount": {
			"type": "string"
		},
		"makerFee": {
			"type": "string"
		},
		"takerFee": {
			"type": "string"
		},
		"expirationTimeSeconds": {
			"type": "string"
		},
		"salt": {
			"type": "string"
		},
		"makerAssetData": {
			"type": "string"
		},
		"takerAssetData": {
			"type": "string"
		},
		"makerFeeAssetData": {
			"type": "string"
		},
		"takerFeeAssetData": {
			"type": "string"
		},
		"exchangeAddress": {
			"type": "string"
		},
		"signature": {
			"type": "string"
		},
		"takerAssetAmountLeft": {
			"type": "string"
		}
	},
	"required": [
		"hash",
		"makerAddress",
		"takerAddress",
		"feeRecipientAddress",
		"senderAddress",
		"makerAssetAmount",
		"takerAssetAmount",
		"makerFee",
		"takerFee",
		"expirationTimeSeconds",
		"salt",
		"makerAssetData",
		"takerAssetData",
		"makerFeeAssetData",
		"takerFeeAssetData",
		"exchangeAddress",
		"signature",
		"takerAssetAmountLeft"
	]
}`)

var rs = &jsonschema.RootSchema{}

func initValidator() {
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}
}