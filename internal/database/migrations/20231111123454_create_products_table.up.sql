-- migrations/20231111123454_create_products_table.up.sql

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
    stock INT,
);

INSERT INTO products (name, price, stock)
VALUES
    ('Product 1', 190990, 100),
    ('Product 2', 29990, 50),
    ('Product 3', 90990, 200),
    ('Product 4', 39990, 30),
    ('Product 5', 49990, 75),
    ('Product 6', 14990, 120),
    ('Product 7', 24990, 60),
    ('Product 8', 34990, 25),
    ('Product 9', 19990, 90),
    ('Product 10', 9990, 150);
