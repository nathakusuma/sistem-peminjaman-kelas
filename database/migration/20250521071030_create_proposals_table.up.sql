CREATE TABLE proposals (
    id UUID PRIMARY KEY,
    proposer_email VARCHAR(320) NOT NULL REFERENCES users (email),
    purpose VARCHAR(50) NOT NULL,
    course VARCHAR(50) NOT NULL,
    class_id VARCHAR(3) NOT NULL,
    lecturer VARCHAR(255) NOT NULL,
    starts_at TIMESTAMP NOT NULL,
    ends_at TIMESTAMP NOT NULL,
    occupancy SMALLINT NOT NULL,
    note VARCHAR(1000),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
