-- +goose Up
INSERT INTO News (Id, Title, Content)
VALUES
    (64, 'Lorem ipsum', 'Dolor sit amet <b>foo</b>'),
    (1, 'first', 'tratata');

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
