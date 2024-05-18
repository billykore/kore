CREATE TABLE IF NOT EXISTS product_inventories
(
    id          SERIAL PRIMARY KEY,
    quantity    INT       NOT NULL,
    description TEXT      NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP NULL,
    deleted_at  TIMESTAMP NULL
);