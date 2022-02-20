
package malgova

import (
	"time"

	"github.com/sivamgr/kstreamdb"
)

// CandleStick struct
type CandleStick struct {
	T time.Time
	O float64
	H float64
	L float64
	C float64
	V uint32
}