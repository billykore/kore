CREATE TABLE IF NOT EXISTS auth_activities
(
    id            SERIAL PRIMARY KEY,
    uuid          VARCHAR(255) UNIQUE NOT NULL,
    username      VARCHAR(255)        NOT NULL,
    login_time    TIMESTAMP           NOT NULL,
    logout_time   TIMESTAMP           NULL,
    is_logged_out BOOL      DEFAULT FALSE,
    created_at    TIMESTAMP DEFAULT NOW(),
    updated_at    TIMESTAMP           NULL,
    deleted_at    TIMESTAMP           NULL
);