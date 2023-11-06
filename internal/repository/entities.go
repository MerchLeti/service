package repository

import "errors"

var ErrNotFound = errors.New("not found")

type Category struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Parent *int64 `json:"parent,omitempty"`
}

type Item struct {
	ID          int64
	Name        string
	Description string
	CategoryID  *int64 `db:"category"`
}

type ItemSummary struct {
	ID         int64
	Name       string
	CategoryID *int64 `db:"category"`
}

type Image struct {
	ID       int64
	ItemID   int64 `db:"item"`
	Position int
	URL      string
}

type Property struct {
	ID     int64
	ItemID int64 `db:"item"`
	Name   string
	Value  string
}

type Type struct {
	ID        string
	ItemID    int64 `db:"item"`
	Name      string
	Price     int
	Available int
}
