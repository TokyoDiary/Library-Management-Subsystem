package malgova

import (
	"fmt"
	"math"
	"sort"

	"gonum.org/v1/gonum/stat"
)

type tradeData struct {
	algoName string
	symbol   string
	orders   []orderEntry
	score    AlgoScore
	trades   []tradeEntry
}

// AlgoScore struct
type AlgoScore struct {
	AlgoName string
	Symbol   string
	// stats and scores
	OrdersCount          int
	TradesCount          int
	TradesWon            int
	TradesLost           int
	WinStreak            int
	LossStreak           int
	NetPnl               float64
	NetPnlPercentAverage float64
	NetPnlPercentStdDev  float64

	SQN float64
}

func (t AlgoScore) String() string {
	return fmt.Sprintf("%12s|%20s|%5d|%4d|%4d:%4d|%3d:%3d| %9.2f |%9.2f|%9.2f| %7.3f", t.AlgoName, t.Symbol, t.OrdersCount, t.TradesCount, t.TradesWon, t.TradesLost, t.WinStreak, t.LossStreak, t.NetPnl, t.NetPnlPercentAverage, t.NetPnlPercentStdDev, t.SQN)
}

type tradeEntry struct {
	orders        int
	buyValue      float64
	sellValue     float64
	pnl           float64
	pnlPercentage float64
}

type algoTradeData struct {
	bySymbolTrades map[string]*tradeData
}

func (a *tradeData) add(t orderEntry) {
	a.orders = append(a.orders, t)
}

// reset score
func (a *tradeData) resetScore() {
	a.trades = make([]tradeEntry, 0)
	a.score = AlgoScore{
		AlgoName: a.algoName,
		Symbol:   a.symbol,
	}
}

func (a *tradeData) consolidateTrades() {
	//sort orders by time
	sort.Slice(a.orders, func(i, j int) bool {
		return a.orders[i].at.Before(a.orders[j].at)
	})

	// consolidate orders into trades
	pos := 0
	openTrade := tradeEntry{}
	for _, o := range a.orders {
		if pos == 0 {
			openTrade.orders = 0
			openTrade.buyValue = 0
			openTrade.sellValue = 0
		}
		pos += o.qty
		if o.qty > 0 {
			openTrade.buyValue = float64(o.qty) * o.price
		} else {
			openTrade.sellValue = -float64(o.qty) * o.price
		}
		openTr