
package malgova

import (
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/sivamgr/kstreamdb"
)

// btDayRunner struct
type btDayRunner struct {
	algos               []reflect.Type
	tickManager         map[string]*btTickManager
	algoRunner          map[string]*btAlgoRunner
	flagSymbolAlgoSetup map[string]bool
	orders              []orderEntry
}

func (bt *btDayRunner) instantiateAllAlgosForSymbol(symbol string) {
	//spawn algos for symbol

	for _, a := range bt.algos {
		pAlgo := newAlgoInstance(a, symbol)
		algoID := pAlgo.ID()
		bt.algoRunner[algoID] = pAlgo
		for _, w := range pAlgo.watch {
			if _, ok := bt.tickManager[w]; !ok {
				bt.tickManager[symbol] = new(btTickManager)
			}
			bt.tickManager[symbol].addObserver(algoID)
		}
	}
}

// worker for concurrent algo execution
func algoRunWorker(wg *sync.WaitGroup, algo *btAlgoRunner, bt *btDayRunner) {
	defer wg.Done()
	algo.run()
}

func (bt *btDayRunner) setup(algos []reflect.Type) {
	bt.algos = algos
	bt.tickManager = make(map[string]*btTickManager)
	bt.algoRunner = make(map[string]*btAlgoRunner)