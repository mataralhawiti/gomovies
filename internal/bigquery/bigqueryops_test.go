// Fix this, test
package bigquery

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
)

func ImportParquet(projectID, datasetID, tableID string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	return nil
}
