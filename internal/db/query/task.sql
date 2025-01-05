-- name: CreateTask :one
INSERT INTO tasks (task_name, category_id, description , timeline_type) VALUES ($1, $2, $3, $4)
RETURNING *;




-- name: UpdateTask :one
UPDATE tasks SET  task_name = $2, category_id=$3, description=$4, timeline_type=$5, priority=$6 WHERE task_id = $1
RETURNING *;


-- name: CompleteTask :one
UPDATE tasks SET  is_completed = $2 WHERE task_id = $1
RETURNING *;


-- name: DeleteTask :exec
DELETE FROM tasks WHERE task_id = $1;


-- name: GetAllTasksInCategory :many
SELECT * FROM tasks
WHERE category_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;