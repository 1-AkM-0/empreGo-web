-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		link TEXT UNIQUE NOT NULL,
    status TEXT DEFAULT "new" CHECK (status IN ("new", "applied", "interview", "approved", "rejected")),
    source TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs;
-- +goose StatementEnd
