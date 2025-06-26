CREATE TABLE access_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    repo_requested TEXT NOT NULL,
    package TEXT,
    result TEXT NOT NULL, -- 'GRANTED', 'DENIED'
    client_ip TEXT,
    timestamp TIMESTAMP DEFAULT now()
);