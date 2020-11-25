
package main

import (
	"fmt"
	"time"

	"github.com/markcheno/go-talib"
	"github.com/sivamgr/kstreamdb"
	"github.com/sivamgr/malgova"
)

// Momento AlgoStrategy
type Momento struct {
	symbol string
	cs1m   *malgova.CandlesData
}

// OnTick Method
func (a *Momento) OnTick(t kstreamdb.TickData, b *malgova.Book) {
	if t.TradingSymbol == a.symbol {
		a.cs1m.Update(t)
	}
}

// OnClose method
func (a *Momento) OnClose(b *malgova.Book) {
	b.Exit()
}

// OnPeriodic method
func (a *Momento) OnPeriodic(t time.Time, b *malgova.Book) {
	if a.cs1m.HasChanged(t) && len(a.cs1m.Close) > 15 {
		ltp := a.cs1m.LTP
		ma1 := talib.Sma(a.cs1m.High, 15)
		ma2 := talib.Ema(a.cs1m.Close, 15)
		ma3 := talib.Sma(a.cs1m.Low, 15)
		if b.IsBookClean() && talib.Crossover(ma2, ma1) {
			//fmt.Printf("[%v] Buy @ %.2f\n", t, ltp)
			quantityToBuy := int(b.Cash / ltp)
			b.Buy(quantityToBuy)
		}
		if b.InPosition() && talib.Crossunder(ma2, ma3) {