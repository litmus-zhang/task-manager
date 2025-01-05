-- name: RegisterUser :one
INSERT INTO users (username, email, password_hash, full_name) VALUES ($1, $2, $3, $4) 
RETURNING *;


-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;
