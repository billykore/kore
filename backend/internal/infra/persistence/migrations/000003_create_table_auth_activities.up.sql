CREATE TABLE IF NOT EXISTS auth_activities
(
    id            VARCHAR(255) UNIQUE PRIMARY KEY NOT NULL,
    username      VARCHAR(255)                    NOT NULL,
    token         VARCHAR(255) UNIQUE             NOT NULL,
    login_time    TIMESTAMP                       NOT NULL,
    logout_time   TIMESTAMP                       NULL,
    is_logged_out BOOL DEFAULT FALSE
);