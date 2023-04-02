
package malgova

import (
	"time"

	"github.com/sivamgr/kstreamdb"
)

// Book struct
type Book struct {
	CashAllocated        float64
	Cash                 float64
	Position             int
	IsMarketOrder        bool
	PendingOrderQuantity int
	PendingOrderPrice    float64
	OrderCount           int
}

// OrderManager Interface
type OrderManager interface {
	PlaceLimitOrder(symbol string, qty int, price float64, a AlgoStrategy)
	PlaceMarketOrder(symbol string, qty int, price float64, a AlgoStrategy)
}

// Engine Interface
type Engine interface {
	RegisterAlgo(algo interface{})
	Run(feed *kstreamdb.DB, oms OrderManager)
	SubscribeChannel(Symbol string, a AlgoStrategy)
}