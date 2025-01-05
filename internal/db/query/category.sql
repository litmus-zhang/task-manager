-- name: CreateCategory :one
INSERT INTO categories (category_name, description, user_id) VALUES ($1, $2, $3) 
RETURNING *;


-- name: GetAllUserCategory :many
SELECT * FROM categories 
WHERE user_id = $1 
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;