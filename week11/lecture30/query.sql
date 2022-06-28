-- Example queries for sqlc
CREATE TABLE top_stories_sqlc (
  	ID BIGSERIAL PRIMARY KEY,
    TITLE TEXT NOT NULL ,
    SCORE INTEGER NOT NULL,
    TIME_OF_SETTING INTEGER NOT NULL
);

-- name: GetRowsCount :one
SELECT COUNT(id) FROM top_stories_sqlc;

-- name: GetLatestTime :one
SELECT MAX(TIME_OF_SETTING) FROM top_stories_sqlc;

-- name: GetDataForStories :many
SELECT TITLE, SCORE FROM top_stories_sqlc;

-- name: DeleteEverything :exec
DELETE FROM top_stories_sqlc;

-- name: InsertIntoDB :one
INSERT INTO top_stories_sqlc (TITLE, SCORE, TIME_OF_SETTING) VALUES ( $1, $2 , $3 ) RETURNING * ;