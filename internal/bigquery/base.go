package bigquery

import (
	"context"
	"fmt"
	"log"

	bq "cloud.google.com/go/bigquery"
	"github.com/mataralhawiti/gomovies/internal/parser" // the directory Not the file!
	"google.golang.org/api/iterator"
)

const BigQuery string = "BigQuery"

func CreateBqClient(projectID string) *bq.Client {
	ctx := context.Background()

	// Creates a client.
	client, err := bq.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	return client
}

func CreateDataSet(datasetName string, bqClinet *bq.Client) {
	ctx := context.Background()
	if err := bqClinet.Dataset(datasetName).Create(ctx, &bq.DatasetMetadata{}); err != nil {
		log.Fatalf("Failed to create dataset: %v", err)
	}
	fmt.Printf("Dataset created\n")
}

func CreateTable(datasetName string, tableName string, bqClinet *bq.Client) {
	schema, err := bq.InferSchema(parser.Movie{})
	if err != nil {
		log.Fatalf("Failed to create dataset: %v", err)
	}
	//schema, err := bqClinet.Dataset(datasetName).Table("ss").

	ctx := context.Background()
	if err := bqClinet.Dataset(datasetName).Table(tableName).Create(ctx, &bq.TableMetadata{Schema: schema}); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func InsertIntoBq(datasetName string, tableName string, bqClinet *bq.Client, mvs []parser.Movie) error {
	ins := bqClinet.Dataset(datasetName).Table(tableName).Inserter()

	ctx := context.Background()
	if err := ins.Put(ctx, mvs); err != nil {
		log.Fatalf("Failed to insert data: %v", err)
		return err
	}
	return nil
}

func ReadFromBq(sqlText string, bqClinet *bq.Client) error {
	q := bqClinet.Query(sqlText)
	fmt.Println(sqlText)
	ctx := context.Background()
	it, err := q.Read(ctx)

	if err != nil {
		log.Fatalf("could not read from BigQuery %v", err)
	}

	for {
		var values []bq.Value //Interface
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf(" %v", err)
		}
		fmt.Println(values)
	}

	return nil
}

func ReadFromBqDryRun(sqlText string, bqClinet *bq.Client) error {
	ctx := context.Background()

	q := bqClinet.Query(sqlText)
	q.DryRun = true

	job, err := q.Run(ctx)
	if err != nil {
		log.Fatalf("could not read from BigQuery %v", err)
	}

	// Dry run is not asynchronous, so get the latest status and statistics.
	status := job.LastStatus()
	if err := status.Err(); err != nil {
		return err
	}

	// fmt.Fprintf(w, "This query will process %d bytes\n", status.Statistics.TotalBytesProcessed)
	fmt.Printf("This query will process %d bytes\n", status.Statistics.TotalBytesProcessed)
	return nil
}
