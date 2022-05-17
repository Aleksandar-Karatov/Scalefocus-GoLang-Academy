package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"
	"week9Lecture26Task/week9/lecture26/dataStructs"
)

type DBcontext struct {
	DB *sql.DB
}

func (dbc *DBcontext) IsDBEmpty() (bool, error) {
	rows, err := dbc.DB.Query("SELECT COUNT(id) FROM TOP_STORIES")
	if err != nil {
		return true, err
	}
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Print(err)
		}
	}
	return count == 0, nil
}

func (dbc *DBcontext) GetLatestTimestamp() (int64, error) {

	if ok, err := dbc.IsDBEmpty(); ok {
		if err != nil {
			return 0, err
		}
		return time.Now().Unix(), nil
	}
	rows, err := dbc.DB.Query("SELECT MAX(TIME_OF_SETTING) FROM TOP_STORIES")
	var lastDBOutTime int64

	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&lastDBOutTime)
		if err != nil {
			return 0, err
		}
	}
	return lastDBOutTime, nil
}
func (dbc *DBcontext) GetAllStoriesAsJSON() ([]byte, error, dataStructs.Stories) {
	if ok, err := dbc.IsDBEmpty(); ok {
		if err != nil {
			return nil, err, dataStructs.Stories{}
		}
	}
	var topStories dataStructs.Stories
	log.Println("Taking data from db")
	rows, err := dbc.DB.Query("SELECT TITLE, SCORE FROM TOP_STORIES")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		cachedStory := dataStructs.Story{}
		err = rows.Scan(&cachedStory.Title, &cachedStory.Score)
		if err != nil {
			log.Println(err)
		}
		topStories.AllStories = append(topStories.AllStories, cachedStory)
	}
	asJSON, errDB := json.Marshal(topStories)
	return asJSON, errDB, topStories

}

func (dbc *DBcontext) removeEverything() error {
	if ok, err := dbc.IsDBEmpty(); ok {
		if err != nil {
			return err
		}
	}

	_, err := dbc.DB.Exec("DELETE FROM TOP_STORIES")
	if err != nil {
		return err
	}
	return nil

}
func (dbc *DBcontext) AddIntoDB(topStories dataStructs.Stories) error {
	err := dbc.removeEverything()
	if err != nil {
		return err
	}
	for _, story := range topStories.AllStories {
		timeInMillis := time.Now().UnixMilli()
		_, err := dbc.DB.Exec("INSERT INTO TOP_STORIES(TITLE, SCORE, TIME_OF_SETTING) values(?,?,?)", story.Title, story.Score, timeInMillis)
		if err != nil {
			return err
		}
	}
	return nil
}
