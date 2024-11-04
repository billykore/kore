CREATE TABLE shipper
(
    id               SERIAL PRIMARY KEY,
    shipper_name     VARCHAR(100) NOT NULL,
    shipping_type    VARCHAR(100) NOT NULL,
    customer_address VARCHAR(100) NOT NULL,
    customer_name    VARCHAR(100) NOT NULL,
    sender_name      VARCHAR(100) NOT NULL,
    fee              BIGINT       NOT NULL,
    status           VARCHAR(100) NOT NULL,
    created_at       TIMESTAMP DEFAULT NOW(),
    updated_at       TIMESTAMP    NULL,
    deleted_at       TIMESTAMP    NULL
);