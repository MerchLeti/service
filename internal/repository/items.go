package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type ItemsRepository struct {
	ds DataSource
}

func Items(ds DataSource) *ItemsRepository {
	return &ItemsRepository{ds: ds}
}

func (r *ItemsRepository) GetItems(ctx context.Context, page, perPage int, categories ...int64) ([]ItemSummary, error) {
	ans := make([]ItemSummary, 0)
	if err := r.ds.Select(
		ctx,
		&ans,
		`select id, name, category from items where category = any($1) order by hits desc limit $2 offset $3`,
		categories,
		perPage,
		(page-1)*perPage,
	); err != nil {
		return nil, fmt.Errorf("couldn't get items: %w", err)
	}
	return ans, nil
}

func (r *ItemsRepository) GetItem(ctx context.Context, id int64) (result Item, err error) {
	err = r.ds.Get(
		ctx,
		&result,
		`update items set hits = hits + 1 where id = $1 returning id, name, category`,
		id,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = ErrNotFound
		}
		err = fmt.Errorf("couldn't get item %d: %w", id, err)
	}
	return
}
