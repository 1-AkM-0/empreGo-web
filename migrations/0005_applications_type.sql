-- +goose Up
-- +goose StatementBegin

ALTER TABLE applications 
  DROP COLUMN title;

ALTER TABLE applications 
  DROP COLUMN source;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE applications;
-- +goose StatementEnd
