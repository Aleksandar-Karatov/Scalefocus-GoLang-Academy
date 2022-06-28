CREATE TABLE users (
  	ID_OF_USER BIGSERIAL PRIMARY KEY,
    USERNAME TEXT NOT NULL ,
    PASSWORD CHAR(60) NOT NULL
);

-- name: GetUserByID :one
SELECT * FROM users WHERE ID_OF_USER = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users WHERE USERNAME = $1 LIMIT 1;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE ID_OF_USER = $1;

-- name: InsertUserInDB :one
INSERT INTO users (USERNAME, PASSWORD) VALUES ( $1, $2 ) RETURNING * ;

CREATE TABLE lists(
	ID_OF_LIST BIGSERIAL PRIMARY KEY,
    NAME TEXT NOT NULL,
  	USERID BIGINT NOT NULL
);

-- name: GetListsForCurrentUser :many
SELECT ID_OF_LIST, NAME FROM lists WHERE USERID = $1;
-- name: GetListById :one
SELECT * FROM lists WHERE ID_OF_LIST = $1 LIMIT 1;
-- name: DeleteListByID :exec
DELETE FROM lists WHERE ID_OF_LIST = $1;
-- name: InsertListInDB :one
INSERT INTO lists (NAME, USERID) VALUES ($1, $2) RETURNING *;

CREATE TABLE tasks(
	ID_OF_TASK INT PRIMARY KEY,
  	TEXT TEXT NOT NULL,
  	LISTID BIGINT NOT NULL,
  	COMPLETED BOOL NOT NULL
);
-- name: GetTasksForCurrentList :many
SELECT * FROM tasks WHERE LISTID = $1;
-- name: DeleteTasktByID :exec
DELETE FROM tasks WHERE ID_OF_TASK = $1;
-- name: InsertTaskInDB :one
INSERT INTO tasks (TEXT, LISTID, COMPLETED) VALUES ($1, $2, $3) RETURNING *;
-- name: PatchTaskInDB :exec
UPDATE tasks set COMPLETED = $1 WHERE ID_OF_TASK = $2;
-- name: GetTaskByID :one
SELECT * FROM tasks WHERE ID_OF_TASK = $1 LIMIT 1;

