CREATE TABLE users (
    email VARCHAR(320) PRIMARY KEY,
    password_hash CHAR(60) NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(7) NOT NULL CHECK ( role IN ('student', 'admin'))
);
