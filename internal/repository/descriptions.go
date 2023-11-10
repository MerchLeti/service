package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DescriptionsRepository struct {
	ds DataSource
}

func Descriptions(ds DataSource) *DescriptionsRepository {
	return &DescriptionsRepository{ds: ds}
}

func (r *DescriptionsRepository) GetAll(ctx context.Context, item int64) ([]Description, error) {
	var ans []Description
	err := r.ds.Select(ctx, &ans, `select * from descriptions where item = $1 order by position`, item)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []Description{}, nil
		}
		return nil, fmt.Errorf("couldn't get descriptions of item %d: %v", item, err)
	}
	return ans, nil
}
