-- Subscription plans table
CREATE TABLE subscription_plans (
    plan_id SERIAL PRIMARY KEY,
    plan_name VARCHAR(50) NOT NULL,
    description TEXT,
    price DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Users table with subscription reference
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    subscription_plan_id INTEGER NOT NULL REFERENCES subscription_plans(plan_id),
    subscription_start_date TIMESTAMP NOT NULL,
    subscription_end_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categories table with color
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    category_name VARCHAR(100) NOT NULL,
    description TEXT,
    color_hex VARCHAR(7) NOT NULL DEFAULT '#808080',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, category_name),
    CHECK (color_hex ~ '^#[0-9A-Fa-f]{6}$')
);

-- Timeline types enum
CREATE TYPE timeline_type AS ENUM ('daily', 'weekly', 'monthly');

-- Task priority enum
CREATE TYPE priority_level AS ENUM ('low', 'medium', 'high', 'urgent');

-- Function to calculate proper due date based on timeline type
CREATE OR REPLACE FUNCTION calculate_due_date(task_timeline timeline_type, created TIMESTAMP)
RETURNS TIMESTAMP AS $$
BEGIN
    CASE task_timeline
        WHEN 'daily' THEN
            -- Due date is start of next day (midnight)
            RETURN DATE_TRUNC('day', created + INTERVAL '1 day');
        WHEN 'weekly' THEN
            -- Due date is start of next week (midnight Monday)
            RETURN DATE_TRUNC('week', created + INTERVAL '1 week') + INTERVAL '1 week';
        WHEN 'monthly' THEN
            -- Due date is start of next month (midnight 1st)
            RETURN DATE_TRUNC('month', created + INTERVAL '1 month');
    END CASE;
END;
$$ LANGUAGE plpgsql IMMUTABLE;

-- Tasks table with automatic due date calculation
CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(category_id),
    task_name VARCHAR(200) NOT NULL,
    description TEXT,
    timeline_type timeline_type NOT NULL,
    priority priority_level NOT NULL DEFAULT 'medium',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL GENERATED ALWAYS AS (calculate_due_date(timeline_type, created_at)) STORED,
    is_completed BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default subscription plan
INSERT INTO subscription_plans (plan_name, description, price)
VALUES ('Basic', 'Basic subscription plan with limited features', 0.00);

-- Create an index on priority for faster task filtering
CREATE INDEX idx_tasks_priority ON tasks(priority);
CREATE INDEX idx_tasks_due_date ON tasks(due_date);

-- -- View for tasks with their due dates
CREATE VIEW task_deadlines AS
SELECT 
    t.task_id,
    t.task_name,
    c.category_name,
    c.color_hex,
    t.timeline_type,
    t.priority,
    t.due_date,
    t.is_completed,
    CASE 
        WHEN t.timeline_type = 'daily' THEN 'Due tomorrow at midnight'
        WHEN t.timeline_type = 'weekly' THEN 'Due next Monday at midnight'
        WHEN t.timeline_type = 'monthly' THEN 'Due first day of next month at midnight'
    END as due_date_description
FROM tasks t
JOIN categories c ON t.category_id = c.category_id
ORDER BY t.due_date;