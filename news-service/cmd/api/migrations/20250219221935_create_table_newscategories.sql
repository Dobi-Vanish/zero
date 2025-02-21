-- +goose Up
    CREATE TABLE IF NOT EXISTS NewsCategories (
        NewsId BIGINT NOT NULL,
        CategoryId BIGINT NOT NULL,
        PRIMARY KEY (NewsId, CategoryId)
    );
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS newscategories;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
