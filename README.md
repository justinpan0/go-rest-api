[![Go Report Card](https://goreportcard.com/badge/github.com/zimengpan/go-rest-api)](https://goreportcard.com/report/github.com/zimengpan/go-rest-api)

This is an implementation of Boomflow cryptocurrency exchange. The project meant to build a Distributed Exchange with Matching strategy. 

As pointed by 0x Documentation:
```
A relayer may chose to accept and broadcast orders where the taker field is equal to an address controlled by the relayer. When the relayer receives two orders on opposite sides of the market with overlapping prices, the relayer can call matchOrders to atomically match both orders
```
**Benefits**
1. Remove the primary race condition in matching proceess from the root, because only the matcher is allowed to fill each order
2. Respond instantly when orders are matched and the orderbook gets upated as soon as the transaction is submitted, allowing for a more real-time trading experience
3. Moves all gas costs to the matcher and potentially provide a better UX for traders

**Trust**

The matcher endpoint is hosted by Conflux Foundation. It adds an element of trust. Conflux Foundation is incentivized to match orders without fees and/or arbitrage opportunities

## <s>Architecture</s>


## Open Source
The structure and codebase of the exchange is forked from GitBitEx, an open source cryptocurrency exchange.

## Dependencies
* <s>MySql (**BINLOG[ROW format]** enabled)</s>
* Kafka
* <s>Redis</s>

## Install
### Server
* git clone the repo
* Run go build
* Run ./go-rest-api