package malgova

import (
	"time"

	"github.com/sivamgr/kstreamdb"
)

//PlaybackFeed struct
type PlaybackFeed struct {
	dateToPlay time.Time
	db         kstreamdb.DB
}

// Run PlaybackFeed
func (f *PlaybackFeed) Run(fCallback func(t kstreamdb.Ti