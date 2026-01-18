package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// stores all the events from polymarket API
//
// we compare the state of this table to the TableEventMarket to see
// if there are new markets
type TableEvent struct {
	EventID   int       `db:"eventID"`
	CreatedAt time.Time `db:"createdAt"`
	Ticker    string    `db:"ticker"`
	ImageURL  *string   `db:"imageURL"`
}

func (t TableEvent) TableName() string {
	return "tbPolymarketEvents"
}

func (db DBManager) UpdateEvent(ctx context.Context, tableEvent TableEvent) error {
	req := fmt.Sprintf(`
        UPDATE %s
        SET createdAt = @createdAt, 
            ticker = @ticker,
            imageURL = @imageURL
		WHERE eventID = @eventID`, tableEvent.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"eventID":   tableEvent.EventID,
		"createdAt": tableEvent.CreatedAt,
		"ticker":    tableEvent.Ticker,
		"imageURL":  tableEvent.ImageURL,
	})
	return err
}

func (db DBManager) InsertEvent(ctx context.Context, tableEvent TableEvent) error {
	req := fmt.Sprintf(`INSERT INTO %s 
	(eventID, createdAt, ticker, imageURL) VALUES 
	(@eventID, @createdAt, @ticker, @imageURL)`, tableEvent.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"eventID":   tableEvent.EventID,
		"createdAt": tableEvent.CreatedAt,
		"ticker":    tableEvent.Ticker,
		"imageURL":  tableEvent.ImageURL,
	})
	return err
}

func (db DBManager) BulkInsertEvent(ctx context.Context, tableEvents []TableEvent) error {
	columns := []string{"eventID", "createdAt", "ticker", "imageURL"}

	_, err := db.Pool.CopyFrom(
		ctx,
		pgx.Identifier{TableEvent{}.TableName()},
		columns,
		pgx.CopyFromSlice(len(tableEvents), func(i int) ([]any, error) {
			return []any{
				tableEvents[i].EventID,
				tableEvents[i].CreatedAt,
				tableEvents[i].Ticker,
				tableEvents[i].ImageURL,
			}, nil
		}),
	)

	if err != nil {
		return fmt.Errorf("couldn't bulk insert into %s table: %w", TableEvent{}.TableName(), err)
	}

	return nil
}

func (db DBManager) GetTableEvents(ctx context.Context, eventid *int, ticker *string) ([]TableEvent, error) {
	// build query
	conditions := []string{}
	args := []any{}
	req := fmt.Sprintf("SELECT * FROM %s", TableEvent{}.TableName())

	if eventid != nil {
		args = append(args, *eventid)
		conditions = append(conditions, fmt.Sprintf("eventID = $%d", len(args)))
	}
	if ticker != nil {
		args = append(args, *ticker)
		conditions = append(conditions, fmt.Sprintf("ticker = $%d", len(args)))
	}
	if len(conditions) > 0 {
		req = fmt.Sprintf("%s WHERE %s", req, strings.Join(conditions, " AND "))
	}

	rows, err := db.Pool.Query(ctx, req, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	output := []TableEvent{}
	for rows.Next() {
		user := TableEvent{}
		err := rows.Scan(
			&user.EventID,
			&user.CreatedAt,
			&user.Ticker,
			&user.ImageURL,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		output = append(output, user)
	}
	return output, nil
}

func (db DBManager) DeleteEvent(ctx context.Context, eventID int) error {
	req := fmt.Sprintf(`
	DELETE FROM %s 
	WHERE eventID = @eventID`, TableEvent{}.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"eventID": eventID,
	})
	return err
}
