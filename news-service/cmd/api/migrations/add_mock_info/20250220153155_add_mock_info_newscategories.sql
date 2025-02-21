-- +goose Up
INSERT INTO NewsCategories (NewsId, CategoryId)
VALUES
    (1, 1),
    (64, 1),
    (64, 2),
    (64, 3);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
