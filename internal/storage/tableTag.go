package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

// stores all the tags from polymarket API (/tags)
type TableTag struct {
	TagID   int    `db:"tagID"`
	Name    string `db:"name"`
	Tracked bool   `db:"tracked"` // whether we are tracking this tag or not
}

func (t TableTag) TableName() string {
	return `"tbPolymarketTags"`
}

func (t TableTag) TableNameIdentifier() string {
	return "tbPolymarketTags"
}

func (db DBManager) UpdateTag(ctx context.Context, tableTag TableTag) error {
	req := fmt.Sprintf(`
        UPDATE %s
        SET name = @name,
            tracked = @tracked
		WHERE "tagID" = @tagID`, tableTag.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"tagID":   tableTag.TagID,
		"name":    tableTag.Name,
		"tracked": tableTag.Tracked,
	})
	return err
}

func (db DBManager) SetTrackedTag(ctx context.Context, tagID int, tracked bool) error {
	req := fmt.Sprintf(`
        UPDATE %s
        SET tracked = @tracked 
		WHERE "tagID" = @tagID`, TableTag{}.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"tagID":   tagID,
		"tracked": tracked,
	})
	return err
}

func (db DBManager) InsertTag(ctx context.Context, tableTag TableTag) error {
	req := fmt.Sprintf(`INSERT INTO %s 
	("tagID", name, tracked) VALUES 
	(@tagID, @name, @tracked)`, tableTag.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"tagID":   tableTag.TagID,
		"name":    tableTag.Name,
		"tracked": tableTag.Tracked,
	})
	return err
}

func (db DBManager) BulkInsertTags(ctx context.Context, tableTags []TableTag) error {
	columns := []string{"tagID", "name", "tracked"}

	_, err := db.Pool.CopyFrom(
		ctx,
		pgx.Identifier{TableTag{}.TableNameIdentifier()},
		columns,
		pgx.CopyFromSlice(len(tableTags), func(i int) ([]any, error) {
			return []any{
				tableTags[i].TagID,
				tableTags[i].Name,
				tableTags[i].Tracked,
			}, nil
		}),
	)

	if err != nil {
		return fmt.Errorf("couldn't bulk insert into %s table: %w", TableTag{}.TableName(), err)
	}

	return nil
}

func (db DBManager) GetTableTags(ctx context.Context, tagID *int, name *string) ([]TableTag, error) {
	// build query
	conditions := []string{}
	args := []any{}
	req := fmt.Sprintf("SELECT * FROM %s", TableTag{}.TableName())

	if tagID != nil {
		args = append(args, *tagID)
		conditions = append(conditions, fmt.Sprintf(`"tagID" = $%d`, len(args)))
	}
	if name != nil {
		args = append(args, *name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(args)))
	}
	if len(conditions) > 0 {
		req = fmt.Sprintf("%s WHERE %s", req, strings.Join(conditions, " AND "))
	}

	rows, err := db.Pool.Query(ctx, req, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	output := []TableTag{}
	for rows.Next() {
		user := TableTag{}
		err := rows.Scan(
			&user.TagID,
			&user.Name,
			&user.Tracked,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		output = append(output, user)
	}
	return output, nil
}

func (db DBManager) GetTrackedTags(ctx context.Context) ([]TableTag, error) {
	req := fmt.Sprintf("SELECT * FROM %s WHERE tracked = TRUE", TableTag{}.TableName())
	rows, err := db.Pool.Query(ctx, req)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[TableTag])
}

func (db DBManager) DeleteTag(ctx context.Context, tagID int) error {
	req := fmt.Sprintf(`DELETE FROM %s WHERE "TagID" = @TagID`, TableTag{}.TableName())

	_, err := db.Pool.Exec(ctx, req, pgx.NamedArgs{
		"tagID": tagID,
	})
	return err
}
