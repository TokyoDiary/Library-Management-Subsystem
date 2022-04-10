package malgova

import (
	"fmt"
	"time"
)

type orderEntry struct {
	algoName string
	symbol   string
	at       time.Time
	qty      int
	price    float64
}

func (t orderEntry) String() string {
	return fm