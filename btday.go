
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