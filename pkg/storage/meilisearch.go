package storage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/meilisearch/meilisearch-go"
)

func NewMeiliClient(host string, port int, indexName, masterKey string) (meilisearch.IndexManager, error) {
	addr := fmt.Sprintf("http://%s:%d", host, port)

	client := meilisearch.New(
		addr,
		meilisearch.WithAPIKey(masterKey),
		meilisearch.WithCustomRetries(
			[]int{http.StatusUnauthorized},
			3,
		),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.CreateIndexWithContext(ctx, &meilisearch.IndexConfig{
		Uid:        indexName,
		PrimaryKey: "id",
	})
	if err != nil {
		return meilisearch.IndexResult{}, fmt.Errorf("failed to create meilisearch index '%s': %w", indexName, err)
	}

	_, err = client.Index(indexName).UpdateFilterableAttributes(&[]string{"rack"})
	if err != nil {
		return meilisearch.IndexResult{}, fmt.Errorf(
			"failed to update meilisearch 'rack' filterable '%s': %w",
			indexName,
			err,
		)
	}

	if !client.IsHealthy() {
		return meilisearch.IndexResult{}, errors.New("meilisearch service is unhealthy")
	}

	return client.Index(indexName), nil
}
