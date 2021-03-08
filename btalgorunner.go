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