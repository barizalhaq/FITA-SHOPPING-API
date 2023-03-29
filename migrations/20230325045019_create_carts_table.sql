-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,

    CONSTRAINT fk_carts_users
        FOREIGN KEY(user_id)
            REFERENCES users(id)
                ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd
