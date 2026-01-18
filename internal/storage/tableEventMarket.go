package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

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

func (db DBManager) UpdateEventMarket(ctx context.Context, tableEventMarket TableEventMarket) error {
	req := fmt.Sprintf(`
        UPDATE %s
        SET eventID = @eventID, 
            question = @question,
            imageURL = @imageURL
		WHERE eventMarketID = @eventMarketID`, tableEventMarket.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"eventMarketID": tableEventMarket.EventMarketID,
		"eventID":       tableEventMarket.EventID,
		"question":      tableEventMarket.Question,
		"imageURL":      tableEventMarket.ImageURL,
	})
	return err
}

func (db DBManager) InsertEventMarket(ctx context.Context, tableEventMarket TableEventMarket) error {
	req := fmt.Sprintf(`INSERT INTO %s 
	(eventMarketID, eventID, question, imageURL) VALUES 
	(@eventMarketID, @eventID, @question, @imageURL)`, tableEventMarket.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"eventMarketID": tableEventMarket.EventMarketID,
		"eventID":       tableEventMarket.EventID,
		"question":      tableEventMarket.Question,
		"imageURL":      tableEventMarket.ImageURL,
	})
	return err
}

func (db DBManager) BulkInsertEventMarket(ctx context.Context, tableEventMarkets []TableEventMarket) error {
	columns := []string{"eventMarketID", "eventID", "question", "imageURL"}

	_, err := db.Pool.CopyFrom(
		ctx,
		pgx.Identifier{TableEventMarket{}.TableName()},
		columns,
		pgx.CopyFromSlice(len(tableEventMarkets), func(i int) ([]any, error) {
			return []any{
				tableEventMarkets[i].EventMarketID,
				tableEventMarkets[i].EventID,
				tableEventMarkets[i].Question,
				tableEventMarkets[i].ImageURL,
			}, nil
		}),
	)

	if err != nil {
		return fmt.Errorf("couldn't bulk insert into %s table: %w", TableEventMarket{}.TableName(), err)
	}

	return nil
}

func (db DBManager) GetTableEventMarkets(ctx context.Context, eventmarketid *int, question *string) ([]TableEventMarket, error) {
	// build query
	conditions := []string{}
	args := []any{}
	req := fmt.Sprintf("SELECT * FROM %s", TableEventMarket{}.TableName())

	if eventmarketid != nil {
		args = append(args, *eventmarketid)
		conditions = append(conditions, fmt.Sprintf("eventMarketID = $%d", len(args)))
	}
	if question != nil {
		args = append(args, *question)
		conditions = append(conditions, fmt.Sprintf("question = $%d", len(args)))
	}
	if len(conditions) > 0 {
		req = fmt.Sprintf("%s WHERE %s", req, strings.Join(conditions, " AND "))
	}

	rows, err := db.Pool.Query(ctx, req, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	output := []TableEventMarket{}
	for rows.Next() {
		user := TableEventMarket{}
		err := rows.Scan(
			&user.EventMarketID,
			&user.EventID,
			&user.Question,
			&user.ImageURL,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		output = append(output, user)
	}
	return output, nil
}

func (db DBManager) DeleteEventMarket(ctx context.Context, eventMarketID int) error {
	req := fmt.Sprintf(`
	DELETE FROM %s 
	WHERE eventMarketID = @eventMarketID`, TableEventMarket{}.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"eventMarketID": eventMarketID,
	})
	return err
}
