-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS status (
    id INTEGER PRIMARY KEY,
    status_name TEXT NOT NULL
);

INSERT INTO status (id, status_name) VALUES (1, "Nova"), (2, "Aplicada"), (3, "Entrevista"), (4, "Aprovado"), (5, "Rejeitado");

CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		link TEXT UNIQUE,
    status_id INTEGER REFERENCES status(id),
    source TEXT NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs;
DROP TABLE status;
-- +goose StatementEnd
