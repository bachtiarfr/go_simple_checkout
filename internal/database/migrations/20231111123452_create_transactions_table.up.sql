-- migrations/20231111123452_create_transactions_table.up.sql

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    total_amount DECIMAL(10, 2) NOT NULL,
    transaction_date TIMESTAMP
);
