// unusual updates to potentially perform manually
package actions

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Eraguzy/PolyNew/internal/storage"
	"github.com/Eraguzy/PolyNew/messager/polymarket"
)

// retrieve and store all sports from polymarket
// (will fail if id already exists)
func UpdateSports(ctx context.Context, db storage.DBManager) error {
	body, err := polymarket.SendGetRequest(
		polymarket.URLGammaAPI,
		polymarket.PathSports,
		nil,
	)
	if err != nil {
		log.Fatalf("error sending get request: %v\n", err)
	}

	sports, storageSports := []polymarket.Sport{}, []storage.TableSport{}
	err = json.Unmarshal(body, &sports)
	if err != nil {
		log.Fatalf("error unmarshaling sports: %v\n", err)
	}
	for _, sport := range sports {
		storageSports = append(storageSports, sport.ToTableSport())
	}

	err = db.BulkInsertSport(ctx, storageSports)
	if err != nil {
		log.Fatalf("error bulk inserting sports: %v\n", err)
	}
	return nil
}
