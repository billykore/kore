CREATE TABLE IF NOT EXISTS products
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255)        NOT NULL,
    description  TEXT                NOT NULL,
    sku          VARCHAR(255) UNIQUE NOT NULL,
    price        BIGINT              NOT NULL DEFAULT 0,
    category_id  INT                 NOT NULL,
    inventory_id INT                 NOT NULL,
    discount_id  INT                 NOT NULL,
    created_at   TIMESTAMP                    DEFAULT NOW(),
    updated_at   TIMESTAMP           NULL,
    deleted_at   TIMESTAMP           NULL
);