package storage

// stores all the events from polymarket API (nested in /events endpoint)
// point is to store all the markets related to an event
type TableEventMarket struct {
	EventMarketID int     `db:"eventMarketID"`
	EventID       int     `db:"eventID"` // FK to TableEvent
	Question      string  `db:"question"`
	ImageURL      *string `db:"imageURL"`
}

func (t TableEventMarket) TableName() string {
	return "tbPolymarketEventMarket"
}
