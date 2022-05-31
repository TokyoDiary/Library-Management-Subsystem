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