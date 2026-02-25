-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS applications (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
    user_id TEXT REFERENCES users(id),
    status TEXT DEFAULT "applied" CHECK ( status IN ("applied", "interview", "approved", "rejected") ) NOT NULL,
    source TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE applications;
-- +goose StatementEnd
