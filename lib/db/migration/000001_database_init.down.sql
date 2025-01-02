-- Drop indexes
DROP INDEX IF EXISTS idx_tasks_priority;
DROP INDEX IF EXISTS idx_tasks_due_date;

-- Drop tasks table
DROP TABLE IF EXISTS tasks;

-- Drop function calculate_due_date
DROP FUNCTION IF EXISTS calculate_due_date;

-- Drop enums
DROP TYPE IF EXISTS priority_level;
DROP TYPE IF EXISTS timeline_type;

-- Drop categories table
DROP TABLE IF EXISTS categories;

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop subscription_plans table
DROP TABLE IF EXISTS subscription_plans;