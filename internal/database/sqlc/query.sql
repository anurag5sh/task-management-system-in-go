-- name: GetUserPassword :one
SELECT password_hash
from users
where username = ? ;

-- name: CreateUser :exec
INSERT INTO users(username,password_hash,email,created_at)
VALUES (?, ?, ?, ?);
