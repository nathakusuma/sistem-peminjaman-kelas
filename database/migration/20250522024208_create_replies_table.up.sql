CREATE TABLE replies (
    id UUID PRIMARY KEY REFERENCES proposals (id),
    admin_id VARCHAR(20) NOT NULL REFERENCES users (id),
    room_id VARCHAR(20) NOT NULL REFERENCES rooms (id),
    is_approved BOOLEAN NOT NULL,
    note VARCHAR(1000),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);
