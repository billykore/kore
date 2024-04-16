CREATE TABLE IF NOT EXISTS product_category
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP    NOT NULL,
    deleted_at  TIMESTAMP    NOT NULL
);