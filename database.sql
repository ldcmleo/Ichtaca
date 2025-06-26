CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    common_name TEXT NOT NULL UNIQUE,
    finger_print TEXT NOT NULL UNIQUE,
    is_admin BOOLEAN DEFAULT FALSE,
    revoked BOOLEAN DEFAULT FALSE
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    scope TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE access_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    repo_requested TEXT NOT NULL,
    package TEXT,
    result TEXT NOT NULL, -- 'GRANTED', 'DENIED'
    client_ip TEXT,
    timestamp TIMESTAMP DEFAULT now()
);