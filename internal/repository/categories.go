package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CategoriesRepository struct {
	ds DataSource
}

func Categories(ds DataSource) *CategoriesRepository {
	return &CategoriesRepository{ds: ds}
}

func (r *CategoriesRepository) getChildren(ctx context.Context, id int64) (ans []Category, err error) {
	err = r.ds.Select(ctx, &ans, `select * from categories where parent = $1`, id)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("couldn't get children of category %d: %w", id, err)
		} else {
			err = nil
		}
	}
	return
}

func (r *CategoriesRepository) GetSubCategories(ctx context.Context, id int64) (ans []Category, err error) {
	if id == 0 {
		return r.GetAll(ctx)
	}
	ans = make([]Category, 0)
	ans = append(ans, Category{})
	err = r.ds.Get(ctx, &ans[0], `SELECT * FROM categories WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("couldn't get category %d: %w", id, err)
	}
	toScan := []int64{ans[0].ID}
	for len(toScan) > 0 {
		nxt := toScan[0]
		toScan = toScan[1:]
		children, err := r.getChildren(ctx, nxt)
		if err != nil {
			return nil, err
		}
		for _, child := range children {
			toScan = append(toScan, child.ID)
			ans = append(ans, child)
		}
	}
	return
}

func (r *CategoriesRepository) GetAll(ctx context.Context) (ans []Category, err error) {
	ans = make([]Category, 0)
	err = r.ds.Select(ctx, &ans, `SELECT * FROM categories`)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ans, nil
		}
		err = fmt.Errorf("couldn't list categories: %w", err)
	}
	return
}
