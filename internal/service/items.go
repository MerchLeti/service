package service

import (
	"context"
	"fmt"
	"math"

	"github.com/MerchLeti/service/internal/repository"
)

const maxItemsPerPage = 100

type itemsRepository interface {
	GetItems(ctx context.Context, page, perPage int, categories ...int64) ([]repository.ItemSummary, error)
	GetItem(ctx context.Context, id int64) (repository.Item, error)
}

type categoriesRepository interface {
	GetSubCategories(ctx context.Context, category int64) ([]repository.Category, error)
}

type typesRepository interface {
	GetAll(ctx context.Context, item int64) ([]repository.Type, error)
}

type imagesRepository interface {
	GetAll(ctx context.Context, item int64) ([]string, error)
	GetAvatar(ctx context.Context, item int64) (string, error)
}

type propertiesRepository interface {
	GetAll(ctx context.Context, item int64) ([]repository.Property, error)
}

type ItemsService struct {
	items      itemsRepository
	categories categoriesRepository
	types      typesRepository
	images     imagesRepository
	properties propertiesRepository
}

func Items(
	items itemsRepository,
	categories categoriesRepository,
	types typesRepository,
	images imagesRepository,
	properties propertiesRepository,
) *ItemsService {
	return &ItemsService{
		items:      items,
		categories: categories,
		types:      types,
		images:     images,
		properties: properties,
	}
}

func (s *ItemsService) GetFromCategory(ctx context.Context, category int64, page, count int) ([]ItemSummary, error) {
	categories, err := s.categories.GetSubCategories(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("couldn't get subcategories of category %d: %w", category, err)
	}
	categoryIDs := make([]int64, len(categories))
	for i := range categories {
		categoryIDs[i] = categories[i].ID
	}
	if count > maxItemsPerPage {
		count = maxItemsPerPage
	}
	items, err := s.items.GetItems(ctx, page, count, categoryIDs...)
	if err != nil {
		return nil, fmt.Errorf("couldn't get items of category %d: %w", category, err)
	}
	ans := make([]ItemSummary, 0, len(items))
	for i := range items {
		price, available, err := s.getItemSummary(ctx, items[i].ID)
		if err != nil {
			return nil, err
		}
		avatar, err := s.images.GetAvatar(ctx, items[i].ID)
		if err != nil {
			return nil, err
		}
		ans = append(ans, NewItemSummary(&items[i], avatar, price, available))
	}
	return ans, nil
}

func (s *ItemsService) GetItem(ctx context.Context, id int64) (item Item, err error) {
	repoItem, err := s.items.GetItem(ctx, id)
	if err != nil {
		return
	}
	types, err := s.GetTypes(ctx, id)
	if err != nil {
		return
	}
	properties, err := s.GetProperties(ctx, id)
	if err != nil {
		return
	}
	images, err := s.GetImages(ctx, id)
	if err != nil {
		return
	}
	item = NewItem(&repoItem, images, properties, types)
	return
}

func (s *ItemsService) GetTypes(ctx context.Context, id int64) (map[string]ItemType, error) {
	types, err := s.types.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}
	ans := make(map[string]ItemType)
	for i := range types {
		ans[types[i].ID] = NewItemType(&types[i])
	}
	return ans, nil
}

func (s *ItemsService) GetProperties(ctx context.Context, id int64) ([]ItemProperty, error) {
	properties, err := s.properties.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}
	ans := make([]ItemProperty, 0, len(properties))
	for i := range properties {
		ans = append(ans, NewItemProperty(&properties[i]))
	}
	return ans, nil
}

func (s *ItemsService) GetImages(ctx context.Context, id int64) ([]string, error) {
	return s.images.GetAll(ctx, id)
}

func (s *ItemsService) getItemSummary(ctx context.Context, item int64) (price, available int, err error) {
	types, err := s.types.GetAll(ctx, item)
	if err != nil {
		return
	}
	if len(types) == 0 {
		return
	}
	price = math.MaxInt
	priceFromAll := math.MaxInt
	for _, t := range types {
		available += t.Available
		if t.Available > 0 && t.Price < price {
			price = t.Price
		}
		if t.Price < priceFromAll {
			priceFromAll = t.Price
		}
	}
	if available == 0 {
		price = priceFromAll
	}
	return
}
