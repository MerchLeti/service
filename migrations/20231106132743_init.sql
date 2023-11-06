-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    parent BIGINT
);
CREATE TABLE items (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    category BIGINT
);
CREATE INDEX items_category ON items(category);
CREATE TABLE images (
    id BIGSERIAL PRIMARY KEY,
    item BIGINT NOT NULL,
    position INT NOT NULL,
    url TEXT NOT NULL
);
CREATE INDEX images_item ON images(item);
CREATE TABLE properties (
    id BIGSERIAL PRIMARY KEY,
    item BIGINT NOT NULL,
    name TEXT NOT NULL,
    value TEXT NOT NULL
);
CREATE INDEX properties_item ON properties(item);
CREATE TABLE types (
    id TEXT NOT NULL,
    item BIGINT NOT NULL,
    name TEXT,
    price INT NOT NULL,
    available INT NOT NULL DEFAULT 0,
    PRIMARY KEY(id, item)
);
CREATE INDEX types_item ON properties(item);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX types_item;
DROP TABLE types;
DROP INDEX properties_item;
DROP TABLE properties;
DROP INDEX images_item;
DROP TABLE images;
DROP INDEX items_category;
DROP TABLE items;
DROP TABLE categories;
-- +goose StatementEnd
