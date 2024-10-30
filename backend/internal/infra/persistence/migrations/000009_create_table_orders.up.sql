CREATE TABLE IF NOT EXISTS orders
(
    id             SERIAL PRIMARY KEY,
    username       VARCHAR(255) NOT NULL,
    cart_ids       INTEGER[]    NOT NULL,
    payment_method VARCHAR(100) NOT NULL,
    status         VARCHAR(100) NOT NULL,
    shipping_id    INT          NULL,
    created_at     TIMESTAMP DEFAULT NOW(),
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL
)