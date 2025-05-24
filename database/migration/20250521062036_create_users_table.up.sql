CREATE TABLE users (
    id VARCHAR(20) PRIMARY KEY,
    email VARCHAR(320) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    role VARCHAR(7) NOT NULL CHECK ( role IN ('student', 'admin'))
);
