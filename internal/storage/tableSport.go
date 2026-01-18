package storage

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

// stores all the sports from polymarket API (/sports)
type TableSport struct {
	SportID   int       `db:"sportID"`
	Sport     string    `db:"sport"`
	CreatedAt time.Time `db:"createdAt"`
	Tracked   bool      `db:"tracked"` // whether we are tracking this sport or not
}

func (t TableSport) TableName() string {
	return "tbPolymarketSport"
}

func (db DBManager) UpdateSport(ctx context.Context, tableSport TableSport) error {
	req := fmt.Sprintf(`
        UPDATE %s
        SET sport = @sport, 
            createdAt = @createdAt, 
            tracked = @tracked 
		WHERE sportID = @sportID`, tableSport.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"sportID":   tableSport.SportID,
		"sport":     tableSport.Sport,
		"createdAt": tableSport.CreatedAt,
		"tracked":   tableSport.Tracked,
	})
	return err
}

func (db DBManager) InsertSport(ctx context.Context, tableSport TableSport) error {
	req := fmt.Sprintf(`INSERT INTO %s 
	(sportID, sport, createdAt, tracked) VALUES 
	(@sportID, @sport, @createdAt, @tracked)`, tableSport.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"sportID":   tableSport.SportID,
		"sport":     tableSport.Sport,
		"createdAt": tableSport.CreatedAt,
		"tracked":   tableSport.Tracked,
	})
	return err
}

func (db DBManager) BulkInsertSport(ctx context.Context, tableSports []TableSport) error {
	columns := []string{"sportID", "sport", "createdAt", "tracked"}

	_, err := db.Pool.CopyFrom(
		ctx,
		pgx.Identifier{TableSport{}.TableName()},
		columns,
		pgx.CopyFromSlice(len(tableSports), func(i int) ([]any, error) {
			return []any{
				tableSports[i].SportID,
				tableSports[i].Sport,
				tableSports[i].CreatedAt,
				tableSports[i].Tracked,
			}, nil
		}),
	)

	if err != nil {
		return fmt.Errorf("couldn't bulk insert into %s table: %w", TableSport{}.TableName(), err)
	}

	return nil
}

func (db DBManager) GetTableSports(ctx context.Context, sportid *int, sport *string) ([]TableSport, error) {
	// build query
	conditions := []string{}
	args := []any{}
	req := fmt.Sprintf("SELECT * FROM %s", TableSport{}.TableName())

	if sportid != nil {
		args = append(args, *sportid)
		conditions = append(conditions, fmt.Sprintf("sportID = $%d", len(args)))
	}
	if sport != nil {
		args = append(args, *sport)
		conditions = append(conditions, fmt.Sprintf("sport = $%d", len(args)))
	}
	if len(conditions) > 0 {
		req = fmt.Sprintf("%s WHERE %s", req, strings.Join(conditions, " AND "))
	}

	rows, err := db.Pool.Query(ctx, req, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	output := []TableSport{}
	for rows.Next() {
		user := TableSport{}
		err := rows.Scan(
			&user.SportID,
			&user.Sport,
			&user.CreatedAt,
			&user.Tracked,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		output = append(output, user)
	}
	return output, nil
}

func (db DBManager) DeleteSport(ctx context.Context, sportID int) error {
	req := fmt.Sprintf(`
	DELETE FROM tbPolymarketSport 
	WHERE sportID = @sportID`, TableSport{}.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"sportID": sportID,
	})
	return err
}
