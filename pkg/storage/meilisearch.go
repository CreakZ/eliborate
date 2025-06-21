package storage

import (
	"context"
	"eliborate/pkg/config"
	"fmt"
	"net/http"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/viper"
)

const MeiliBookIndex = "book_docs"

func NewMeiliClient() meilisearch.IndexManager {
	addr := fmt.Sprintf(
		"http://%s:%d",
		viper.GetString(config.MeiliHost),
		viper.GetInt(config.MeiliPort),
	)

	client := meilisearch.New(
		addr,
		meilisearch.WithAPIKey(viper.GetString(config.MeiliMasterKey)),
		meilisearch.WithCustomRetries(
			[]int{http.StatusUnauthorized},
			3,
		),
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3*time.Second))
	defer cancel()

	_, err := client.CreateIndexWithContext(ctx, &meilisearch.IndexConfig{
		Uid:        MeiliBookIndex,
		PrimaryKey: "id",
	})
	if err != nil {
		panic(fmt.Errorf("meilisearch: failed to create index '%s': %w", MeiliBookIndex, err))
	}

	_, err = client.Index(MeiliBookIndex).UpdateFilterableAttributes(&[]string{"rack"})
	if err != nil {
		panic(fmt.Errorf(
			"meilisearch: failed to update 'rack' filterable '%s': %w",
			MeiliBookIndex,
			err,
		))
	}

	if !client.IsHealthy() {
		panic("meilisearch: server is unhealthy")
	}

	return client.Index(MeiliBookIndex)
}
