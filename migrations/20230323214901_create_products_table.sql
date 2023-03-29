-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    sku VARCHAR UNIQUE NOT NULL,
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL CHECK (price > 0),
    qty INT DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
