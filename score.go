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
	NetPnl  