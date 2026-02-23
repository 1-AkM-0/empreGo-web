-- +goose Up
-- +goose StatementBegin

DROP TABLE users;

CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
	  email TEXT,
    username TEXT NOT NULL,
    github_id TEXT UNIQUE NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
