CREATE TABLE IF NOT EXISTS otp
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(255) NOT NULL,
    otp        VARCHAR(255) NOT NULL,
    is_active  BOOL      DEFAULT TRUE,
    expires_at TIMESTAMP    NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP    NULL,
    deleted_at TIMESTAMP    NULL
);
