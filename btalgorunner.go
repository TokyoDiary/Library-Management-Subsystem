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
	