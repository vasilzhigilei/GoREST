CREATE TABLE IF NOT EXISTS customers(
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
password VARCHAR(100) NOT NULL,
certificates jsonb NOT NULL
);