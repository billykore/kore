CREATE TABLE IF NOT EXISTS discounts
(
    id                  SERIAL PRIMARY KEY,
    name                VARCHAR(255) NOT NULL,
    description         TEXT         NOT NULL,
    discount_percentage DECIMAL      NOT NULL,
    active              BOOL      DEFAULT FALSE,
    created_at          TIMESTAMP DEFAULT NOW(),
    updated_at          TIMESTAMP    NOT NULL,
    deleted_at          TIMESTAMP    NOT NULL
);