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
	pnl           floa