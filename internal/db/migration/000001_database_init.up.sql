

-- Users table with subscription reference
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categories table with color
CREATE TABLE IF NOT EXISTS categories (
    category_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    category_name VARCHAR(100) NOT NULL,
    description TEXT,
    color_hex VARCHAR(7) NOT NULL DEFAULT '#808080',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, category_name),
    CHECK (color_hex ~ '^#[0-9A-Fa-f]{6}$')
);



-- Function to calculate proper due date based on timeline type
CREATE OR REPLACE FUNCTION calculate_due_date(task_timeline INTERVAL, created TIMESTAMP)
RETURNS TIMESTAMP AS $$
BEGIN
    RETURN created + task_timeline;
END;
$$ LANGUAGE plpgsql IMMUTABLE;



-- Tasks table with automatic due date calculation
CREATE TABLE IF NOT EXISTS tasks (
    task_id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(category_id),
    task_name VARCHAR(200) NOT NULL,
    description TEXT,
    timeline_type INTERVAL NOT NULL DEFAULT '1 day',
    priority VARCHAR(200) NOT NULL DEFAULT 'medium',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL GENERATED ALWAYS AS (calculate_due_date(timeline_type, created_at)) STORED,
    is_completed BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



