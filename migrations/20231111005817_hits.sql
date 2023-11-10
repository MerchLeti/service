-- +goose Up
-- +goose StatementBegin
ALTER TABLE items ADD COLUMN hits BIGINT DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE items DROP COLUMN hits;
-- +goose StatementEnd
