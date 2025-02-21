-- +goose Up
    CREATE TABLE IF NOT EXISTS News (
        Id INTEGER PRIMARY KEY,
        Title TEXT NOT NULL,
        Content TEXT NOT NULL
    );
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS news;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
