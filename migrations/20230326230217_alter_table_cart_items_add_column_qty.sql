-- +goose Up
-- +goose StatementBegin
ALTER TABLE cart_items
    ADD COLUMN qty INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE cart_items
    DROP COLUMN qty;
-- +goose StatementEnd
