-- +goose Up
ALTER TABLE NewsCategories
    ADD CONSTRAINT fk_news
        FOREIGN KEY (NewsId)
            REFERENCES News(Id)
            ON DELETE CASCADE;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE NewsCategories
DROP CONSTRAINT IF EXISTS fk_news;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
