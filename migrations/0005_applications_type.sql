-- +goose Up
-- +goose StatementBegin

ALTER TABLE applications 
  DROP COLUMN title;

ALTER TABLE applications 
  DROP COLUMN source;

ALTER TABLE applications
  ADD COLUMN job_id INT REFERENCES jobs(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE applications;
-- +goose StatementEnd
