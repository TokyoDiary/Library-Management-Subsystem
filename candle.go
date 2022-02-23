
package malgova

import (
	"time"

	"github.com/sivamgr/kstreamdb"
)

// CandleStick struct
type CandleStick struct {
	T time.Time
	O float64
	H float64
	L float64
	C float64
	V uint32
}

// CandlesData struct
type CandlesData struct {
	marketStartTimeHour   int
	marketStartTimeMinute int

	dayStartTime time.Time
	candlePeriod int

	currentCandle              CandleStick
	currentCandleHarvestTime   time.Time
	currentCandleTicksReceived uint32

	totalVolumeTraded uint32
	Candles           []CandleStick
	Open              []float64
	High              []float64
	Low               []float64
	Close             []float64
	Volume            []float64
	LTP               float64
}

// HasChanged method, call with timestamp, to harvest candless
// return true is a new candle is formed
func (f *CandlesData) HasChanged(t time.Time) bool {
	if (!t.Before(f.currentCandleHarvestTime)) && f.currentCandleTicksReceived > 0 {
		f.Candles = append(f.Candles, f.currentCandle)
		f.Open = append(f.Open, f.currentCandle.O)
		f.High = append(f.High, f.currentCandle.H)