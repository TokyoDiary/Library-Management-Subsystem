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
	price    floa