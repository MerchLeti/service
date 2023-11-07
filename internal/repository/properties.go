package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PropertiesRepository struct {
	ds DataSource
}

func Properties(ds DataSource) *PropertiesRepository {
	return &PropertiesRepository{ds: ds}
}

func (r *PropertiesRepository) GetAll(ctx context.Context, item int64) ([]Property, error) {
	var properties []Property
	err := r.ds.Select(ctx, &properties, `select * from properties where item = $1 order by position`, item)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []Property{}, nil
		}
		return nil, fmt.Errorf("couldn't get properties of item %d: %v", item, err)
	}
	return properties, nil
}
