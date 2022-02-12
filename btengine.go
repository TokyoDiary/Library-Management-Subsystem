
package malgova

import (
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/sivamgr/kstreamdb"
)

// BacktestEngine struct
type BacktestEngine struct {
	algos  []reflect.Type
	orders []orderEntry
	scores []AlgoScore
}

// RegisterAlgo BacktestEngine
func (bt *BacktestEngine) RegisterAlgo(a interface{}) {
	if bt.algos == nil {
		bt.algos = make([]reflect.Type, 0)
	}
	bt.algos = append(bt.algos, reflect.TypeOf(a))
}

// RunAlgoBetweenDate method
func (bt *BacktestEngine) RunAlgoBetweenDate(feed *kstreamdb.DB, oms OrderManager, algoName string, startDate time.Time, endDate time.Time) {
	selectedAlgo := make([]reflect.Type, 0)
	for _, a := range bt.algos {
		if a.Name() == algoName {
			selectedAlgo = append(selectedAlgo, a)
			break
		}
	}

	if len(selectedAlgo) == 0 {
		return
	}

	dates, _ := feed.GetDates()
	dayRunner := btDayRunner{}
	dayRunner.setup(selectedAlgo)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)