CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(15) NOT NULL,
    account_password BYTEA NOT NULL
);