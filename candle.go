
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
		f.Low = append(f.Low, f.currentCandle.L)
		f.Close = append(f.Close, f.currentCandle.C)
		f.Volume = append(f.Volume, float64(f.currentCandle.V))
		f.currentCandleTicksReceived = 0 // reset candles
		return true
	}
	return false
}

// Update method
func (f *CandlesData) Update(t kstreamdb.TickData) {
	var volume uint32
	var ltt time.Time

	if t.IsTradable {
		ltt = t.LastTradeTime
		volume = t.VolumeTraded
	} else {
		ltt = t.Timestamp
		volume = 0
	}

	f.LTP = float64(t.LastPrice)

	if f.currentCandleTicksReceived == 0 {
		if len(f.Candles) == 0 {
			f.dayStartTime = time.Date(ltt.Year(), ltt.Month(), ltt.Day(), f.marketStartTimeHour, f.marketStartTimeMinute, 0, 0, ltt.Location())
		}
		if ltt.Before(f.dayStartTime) {
			return
		}

		nw := int(ltt.Sub(f.dayStartTime).Seconds()) / f.candlePeriod
		f.currentCandle.T = f.dayStartTime.Add(time.Second * time.Duration(nw*f.candlePeriod))
		f.currentCandleHarvestTime = f.currentCandle.T.Add(time.Second * time.Duration(f.candlePeriod))
		f.currentCandle.O = f.LTP
		f.currentCandle.H = f.LTP