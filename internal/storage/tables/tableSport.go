package storage

import "time"

// stores all the sports from polymarket API (/sports)
type TableSport struct {
	SportID   int       `db:"sportID"`
	Sport     string    `db:"sport"`
	CreatedAt time.Time `db:"createdAt"`
	Tracked   bool      `db:"tracked"` // whether we are tracking this sport or not
}

func setTracked(sportID int, tracked bool) error { return nil }
