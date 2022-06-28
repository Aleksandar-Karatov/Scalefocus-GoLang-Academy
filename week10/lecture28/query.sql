-- name GetRowsCount :one
SELECT COUNT(id) FROM top_stories_sqlc;

--name GetLatestTime :one
SELECT MAX(TIME_OF_SETTING) FROM top_stories_sqlc;

--name GetDataForStories :many
SELECT TITLE, SCORE FROM top_stories_sqlc;

--name DeleteEverything :exec
DELETE FROM top_stories_sqlc;

--name InsertIntoDB :one
INSERT INTO top_stories_sqlc (TITLE, SCORE, TIME_OF_SETTING) VALUES (?,?,?);


