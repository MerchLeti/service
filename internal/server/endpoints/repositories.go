package endpoints

import (
	"context"

	"github.com/MerchLeti/service/internal/repository"
	"github.com/MerchLeti/service/internal/service"
)

type categoriesRepo interface {
	GetSubCategories(ctx context.Context, id int64) (ans []repository.Category, err error)
	GetAll(ctx context.Context) (ans []repository.Category, err error)
}

type itemsService interface {
	GetFromCategory(ctx context.Context, category int64, page, count int) ([]service.ItemSummary, error)
	GetItem(ctx context.Context, id int64) (item service.Item, err error)
}
