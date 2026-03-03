-- +goose Up
-- +goose StatementBegin

ALTER TABLE jobs
  ADD COLUMN company TEXT NOT NULL

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs;
-- +goose StatementEnd
