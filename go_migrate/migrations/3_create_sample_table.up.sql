CREATE TABLE IF NOT EXISTS sample(
 id SERIAL PRIMARY KEY,
 sample_variance VARCHAR(255) NOT NULL,
 label TEXT NOT NULL,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);