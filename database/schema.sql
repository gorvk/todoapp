CREATE TABLE IF NOT EXISTS Todos (
    id UUID PRIMARY KEY,
    title VARCHAR NOT NULL,
    is_completed BOOLEAN NOT NULL,
    user_id VARCHAR NOT NULL
);