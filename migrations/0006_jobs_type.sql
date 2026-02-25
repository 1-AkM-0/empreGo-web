-- +goose Up
-- +goose StatementBegin

ALTER TABLE jobs
  ADD COLUMN type TEXT CHECK (type IN ("backend", "fullstack", "frontend", "geral"));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs;
-- +goose StatementEnd
