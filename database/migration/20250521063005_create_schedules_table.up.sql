CREATE TABLE schedules (
    id UUID PRIMARY KEY,
    day VARCHAR(6) NOT NULL,
    start_time TIME NOT NULL,
    finish_time TIME NOT NULL,
    room_id VARCHAR(20) NOT NULL REFERENCES rooms (id),
    course VARCHAR(50) NOT NULL,
    class_id VARCHAR(3) NOT NULL,
    is_laboratory BOOLEAN NOT NULL,
    lecturer VARCHAR(255) NOT NULL,
    major VARCHAR(30) NOT NULL,
    start_date DATE NOT NULL,
    finish_date DATE NOT NULL
);
