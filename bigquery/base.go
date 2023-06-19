package bigquery

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
)

const BigQuery string = "BigQuery"

func CreateBqClient(projectID string) *bigquery.Client {
	ctx := context.Background()

	// Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	return client
}

func CreateDataSet(datasetName string, bqClinet *bigquery.Client) {
	ctx := context.Background()
	if err := bqClinet.Dataset(datasetName).Create(ctx, &bigquery.DatasetMetadata{}); err != nil {
		log.Fatalf("Failed to create dataset: %v", err)
	}
	fmt.Printf("Dataset created\n")
}

func InsertIntoBq() string {
	return BigQuery
}

func ReadFromBq() string {
	return BigQuery
}
