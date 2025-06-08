CREATE TABLE replies (
    id UUID PRIMARY KEY REFERENCES proposals (id),
    admin_email VARCHAR(320) NOT NULL REFERENCES users (email),
    room VARCHAR(50),
    is_approved BOOLEAN NOT NULL,
    note VARCHAR(1000),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
