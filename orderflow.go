
package malgova

import (
	"fmt"

	"github.com/sivamgr/kstreamdb"
)

// PriceCell struct
type PriceCell struct {
	BidQuantityTaken uint32
	AskQuantityTaken uint32
	VolumeTraded     uint32
}

// OrderFlowMonitor struct
type OrderFlowMonitor struct {
	LastTick               kstreamdb.TickData
	TotalBidsQuantityTaken uint32
	TotalAsksQuantityTaken uint32
	TicksUpdated           uint32
	Prices                 map[uint32]*PriceCell
	Bids                   [5]kstreamdb.DepthItem
	Asks                   [5]kstreamdb.DepthItem
}

// NewOrderFlowMonitor creates a new orderflow monitor
func NewOrderFlowMonitor() *OrderFlowMonitor {
	pNew := new(OrderFlowMonitor)
	pNew.Prices = make(map[uint32]*PriceCell)
	return pNew
}

// GetPriceCell returns the price cell
func (r *OrderFlowMonitor) GetPriceCell(p float32) *PriceCell {
	k := uint32(p * 100)
	if _, ok := r.Prices[k]; !ok {
		r.Prices[k] = new(PriceCell)
	}
	return r.Prices[k]
}

// Update processes the tick
func (r *OrderFlowMonitor) Update(t kstreamdb.TickData) {
	if t.VolumeTraded > r.LastTick.VolumeTraded {