-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
	  email TEXT UNIQUE NOT NULL,
    hash_password TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
