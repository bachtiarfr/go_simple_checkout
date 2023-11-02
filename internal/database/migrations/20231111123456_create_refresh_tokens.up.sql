-- migrations/20231111123456_create_refresh_tokens.up.sql

CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    refresh_token TEXT NOT NULL,
    expiration TIMESTAMP NOT NULL
);