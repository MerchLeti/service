package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type TypesRepository struct {
	ds DataSource
}

func Types(ds DataSource) *TypesRepository {
	return &TypesRepository{ds: ds}
}

func (r *TypesRepository) GetAll(ctx context.Context, item int64) ([]Type, error) {
	var types []Type
	err := r.ds.Select(ctx, &types, `select * from types where item = $1 order by position`, item)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []Type{}, nil
		}
		return nil, fmt.Errorf("couldn't get types of item %d: %v", item, err)
	}
	return types, nil
}
