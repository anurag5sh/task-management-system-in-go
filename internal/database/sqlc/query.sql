-- name: GetUserPassword :one
SELECT id, password_hash
from users
where username = ? ;

-- name: CreateUser :exec
INSERT INTO users(username,password_hash,email,created_at)
VALUES (?, ?, ?, ?);

-- name: GetAllTasks :many
SELECT * FROM tasks where user_id = ?;

-- name: GetTask :one
SELECT * FROM tasks where id = ? and user_id = ?;

-- name: CreateTask :exec
INSERT INTO tasks(title, description, status, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateTask :exec
UPDATE tasks SET title = ?, description = ?, status = ?, updated_at = ? where id = ?;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = ?;