package malgova

import (
	"reflect"
	"time"

	"github.com/sivamgr/kstreamdb"
)

type btAlgoRunner struct {
	algoName            string
	symbol              string
	ptr                 reflect.Value
	ainterface          interface{}
	strategy            AlgoStrategy
	book                Book
	watch               []string
	enable              bool
	lastTick            kstreamdb.TickData
	queueTick           []kstreamdb.TickData
	utcLastPeriodicCall int64
	orders              []orderEntry
}

func (a *btAlgoRunner) ID() string {
	return a.algoName + "::" + a.symbol
}

func (a *btAlgoRunner) queue(t kstreamdb.TickData) {
	if a.enable {
		a.queueTick = append(a.queueTick, t)
	}
}

func (a *btAlgoRunner) resetQueue() {
	a.queueTick = make([]kstreamdb.TickData, 0, len(a.watch)*24000)
}

func (a *btAlgoRunner) run() {
	if a.enable {
		a.strategy.OnDayStart(&a.book)
		for _, t := range a.queueTick {
			a.checkClock(t.Timestamp)
			a.handleTick(t)
		}
		a.strategy.OnDayEnd(&a.book)
		a.resetQueue()
		//fmt.Printf("P/L %9.2f | Trades %3d | %s\n", a.book.Cash-a.book.CashAllocated, a.book.OrderCount, a.ID())
	}
}

func (a *btAlgoRunner) exit() {
	if a.enable {
		a.strategy.OnClose(&a.book)
		a.handleBook()
	}
}

func (a *btAlgoRunner) checkClock(t time.Time) {
	utcNow := t.Unix()
	if a.utcLastPeriodicCall < utcNow {
		a.utcLastPeriodicCall = utcNow
		a.strategy.OnPeriodic(time.Unix(utcNow, 0), &a.book)
	}
}

func (a *btAlgoRunner) handleBook() {
	if a.book.IsOrderWaiting() {
		if a.book.IsMarketOrder {
			if a.book.PendingOrderQuantity > 0 {
				buyPrice := a.lastTick.Ask[0].Price
				if buyPrice <= 0 {
					buyPrice = a.lastTick.LastPrice
				}
				cost := buyPrice * float32(a.book.PendingOrderQuantity)
				a.book.Cash -= float64(cost)
				a.book.Position += a.book.PendingOrderQuantity
				// add trade trade ledger
				a.orders = append(a.orders, or