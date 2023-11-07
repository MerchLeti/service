package service

import "github.com/MerchLeti/service/internal/repository"

type ItemSummary struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CategoryID *int64 `json:"category,omitempty"`
	Avatar     string `json:"avatar"`
	Price      int    `json:"price"`
	Available  int    `json:"available"`
}

type ItemType struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Available int    `json:"available"`
}

type ItemProperty struct {
	Subtitle *string `json:"subtitle,omitempty"`
	Name     string  `json:"name"`
	Value    string  `json:"value"`
}

type Item struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	CategoryID *int64         `json:"category,omitempty"`
	Images     []string       `json:"images"`
	Properties []ItemProperty `json:"properties"`
	Types      []ItemType     `json:"types"`
}

func NewItemSummary(from *repository.ItemSummary, avatar string, price, available int) ItemSummary {
	return ItemSummary{
		ID:         from.ID,
		Name:       from.Name,
		CategoryID: from.CategoryID,
		Avatar:     avatar,
		Price:      price,
		Available:  available,
	}
}

func NewItemType(from *repository.Type) ItemType {
	return ItemType{
		ID:        from.ID,
		Name:      from.Name,
		Price:     from.Price,
		Available: from.Available,
	}
}

func NewItemProperty(from *repository.Property) ItemProperty {
	return ItemProperty{
		Subtitle: from.Subtitle,
		Name:     from.Name,
		Value:    from.Value,
	}
}

func NewItem(from *repository.Item, images []string, properties []ItemProperty, types []ItemType) Item {
	return Item{
		ID:         from.ID,
		Name:       from.Name,
		CategoryID: from.CategoryID,
		Images:     images,
		Properties: properties,
		Types:      types,
	}
}
