package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type ImagesRepository struct {
	ds DataSource
}

func Images(ds DataSource) *ImagesRepository {
	return &ImagesRepository{ds: ds}
}

func (r *ImagesRepository) GetAll(ctx context.Context, item int64) ([]string, error) {
	var images []Image
	err := r.ds.Select(ctx, &images, `select url from images where item = $1 order by position`, item)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("couldn't get images of item %d: %v", item, err)
	}
	ans := make([]string, len(images))
	for i := range images {
		ans[i] = images[i].URL
	}
	return ans, nil
}

func (r *ImagesRepository) GetAvatar(ctx context.Context, item int64) (string, error) {
	var image Image
	err := r.ds.Get(ctx, &image, `SELECT url FROM images WHERE item = $1 ORDER BY position LIMIT 1`, item)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("couldn't get avatar of item %d: %v", item, err)
	}
	return image.URL, nil
}
