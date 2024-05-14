CREATE TABLE IF NOT EXISTS carts
(
    id         SERIAL PRIMARY KEY,
    user_id    INT       NOT NULL,
    product_id INT       NOT NULL,
    quantity   INT       NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);