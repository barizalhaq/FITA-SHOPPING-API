-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cart_items (
    cart_id INT NOT NULL,
    product_id INT NOT NULL,

    CONSTRAINT fk_cart_items_carts
        FOREIGN KEY(cart_id)
            REFERENCES carts(id)
                ON UPDATE CASCADE ON DELETE CASCADE,
    
    CONSTRAINT fk_cart_items_products
        FOREIGN KEY(product_id)
            REFERENCES products(id)
                ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart_items;
-- +goose StatementEnd
