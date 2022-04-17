
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