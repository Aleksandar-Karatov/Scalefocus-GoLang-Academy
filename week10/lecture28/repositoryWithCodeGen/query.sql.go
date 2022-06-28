package repositoryWithCodeGen

import (
	"context"
)

const deleteEverything = `-- name: DeleteEverything :exec
DELETE FROM top_stories_sqlc
`

func (q *Queries) DeleteEverything(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteEverything)
	return err
}

const getDataForStories = `-- name: GetDataForStories :many
SELECT TITLE, SCORE FROM top_stories_sqlc
`

type DataPayloadForRows struct {
	AllStoriesInRows []GetDataForStoriesRow `json:"top_stories"`
}

type GetDataForStoriesRow struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}

type StoriesPageDataFromDB struct {
	PageTitle string
	Links     []GetDataForStoriesRow
}

func (q *Queries) GetDataForStories(ctx context.Context) (DataPayloadForRows, error) {
	rows, err := q.db.QueryContext(ctx, getDataForStories)
	if err != nil {
		return DataPayloadForRows{nil}, err
	}
	defer rows.Close()
	var items DataPayloadForRows
	for rows.Next() {
		var i GetDataForStoriesRow
		if err := rows.Scan(&i.Title, &i.Score); err != nil {
			return DataPayloadForRows{nil}, err
		}
		items.AllStoriesInRows = append(items.AllStoriesInRows, i)
	}
	if err := rows.Close(); err != nil {
		return DataPayloadForRows{nil}, err
	}
	if err := rows.Err(); err != nil {
		return DataPayloadForRows{nil}, err
	}
	return items, nil
}

const getLatestTime = `-- name: GetLatestTime :one
SELECT MAX(TIME_OF_SETTING) FROM top_stories_sqlc
`

func (q *Queries) GetLatestTime(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLatestTime)
	var max int64
	err := row.Scan(&max)
	return max, err
}

const getRowsCount = `-- name: GetRowsCount :one
SELECT COUNT(id) FROM top_stories_sqlc
`

func (q *Queries) GetRowsCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getRowsCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const insertIntoDB = `-- name: InsertIntoDB :one
INSERT INTO top_stories_sqlc (TITLE, SCORE, TIME_OF_SETTING) VALUES ( $1, $2 , $3 ) RETURNING id, title, score, time_of_setting
`

type InsertIntoDBParams struct {
	Title         string
	Score         int
	TimeOfSetting int64
}

func (q *Queries) InsertIntoDB(ctx context.Context, arg InsertIntoDBParams) (TopStoriesSqlc, error) {
	row := q.db.QueryRowContext(ctx, insertIntoDB, arg.Title, arg.Score, arg.TimeOfSetting)
	var i TopStoriesSqlc
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Score,
		&i.TimeOfSetting,
	)
	return i, err
}
