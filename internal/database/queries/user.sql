-- name: GetUserById :one
select * from users u where u.id = $1;