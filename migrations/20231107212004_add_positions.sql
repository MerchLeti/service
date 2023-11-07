-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE position_seq START 1;
ALTER TABLE properties ADD COLUMN position INT DEFAULT nextval('position_seq');
ALTER TABLE types ADD COLUMN position INT DEFAULT nextval('position_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE types DROP COLUMN position;
ALTER TABLE properties DROP COLUMN position;
DROP SEQUENCE position_seq;
-- +goose StatementEnd
