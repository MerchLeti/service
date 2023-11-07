-- +goose Up
-- +goose StatementBegin
ALTER TABLE properties ADD COLUMN subtitle TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE properties DROP COLUMN subtitle;
-- +goose StatementEnd
