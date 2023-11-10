-- +goose Up
-- +goose StatementBegin
CREATE TABLE descriptions (
    id BIGSERIAL PRIMARY KEY,
    item BIGINT NOT NULL,
    title TEXT,
    value TEXT,
    position INT DEFAULT nextval('position_seq')
);
INSERT INTO descriptions (item, value) SELECT id, description FROM items;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE descriptions;
-- +goose StatementEnd
