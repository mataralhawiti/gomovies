package bigquery

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/mataralhawiti/gomovies/parser" // the directory Not the file!
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

func CreateTable(datasetName string, tableName string, bqClinet *bigquery.Client) {
	schema, err := bigquery.InferSchema(parser.Movie{})
	if err != nil {
		log.Fatalf("Failed to create dataset: %v", err)
	}
	//schema, err := bqClinet.Dataset(datasetName).Table("ss").

	ctx := context.Background()
	if err := bqClinet.Dataset(datasetName).Table(tableName).Create(ctx, &bigquery.TableMetadata{Schema: schema}); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func InsertIntoBq(datasetName string, tableName string, bqClinet *bigquery.Client, mvs []parser.Movie) error {
	ins := bqClinet.Dataset(datasetName).Table(tableName).Inserter()

	ctx := context.Background()
	if err := ins.Put(ctx, mvs); err != nil {
		log.Fatalf("Failed to insert data: %v", err)
		return err
	}
	return nil
}

func ReadFromBq() string {
	return BigQuery
}
