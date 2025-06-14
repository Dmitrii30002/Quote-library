DROP TABLE IF EXISTS Quote;
DROP TABLE IF EXISTS Author;

CREATE TABLE Author (
    id SERIAL PRIMARY KEY,
    author_name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);