CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    common_name TEXT NOT NULL UNIQUE,
    finger_print TEXT NOT NULL UNIQUE,
    is_admin BOOLEAN DEFAULT FALSE,
    revoked BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now()
);