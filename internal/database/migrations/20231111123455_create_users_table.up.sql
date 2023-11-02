-- migrations/20231111123455_create_users_table.up.sql

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(30) NOT NULL,
    email TEXT NOT NULL,
    age INT,
    password TEXT NOT NULL,
    phone_number VARCHAR(15)
);

-- Insert dummy user data
INSERT INTO users (fullname, email, age, password, phone_number) VALUES
    ('Test User', 'test@example.com', 35, '$2a$10$GSZl8FNm/DtZ9van76NBK.L4rGjufy00rNHkgTLWGpyouDbteXanG', '+62895334568841');
