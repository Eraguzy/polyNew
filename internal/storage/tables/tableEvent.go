package storage

// stores all the events from polymarket API
//
// we compare the state of this table to the TableEventMarket to see
// if there are new markets
type TableEvent struct {
	EventID   int     `db:"eventID"`
	CreatedAt int     `db:"createdAt"`
	Ticker    string  `db:"ticker"`
	ImageURL  *string `db:"imageURL"`
}
