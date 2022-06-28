package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "SFDB.db")
	mux := http.NewServeMux()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	mux.Handle("/api/top", HandleTop(db))
	log.Fatal(http.ListenAndServe(":9000", mux))
}

type NewsID = int

type IDPayload struct {
	IDs []NewsID
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}
type Stories struct {
	AllStories []Story `json:"top_stories"`
}

type SimpleHandler struct{}

type StoriesPageData struct {
	PageTitle string
	Links     []Story
}

func HandleTop(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var lastDBOutTime int64 // check if there is any data in the db
		rows, err := db.Query("SELECT COUNT(id) FROM TOP_STORIES")
		if err != nil {
			log.Println(err)
		}
		var count int
		for rows.Next() {
			err = rows.Scan(&count)
			if err != nil {
				log.Print(err)
			}
		}

		if count > 0 {
			rows, err := db.Query("SELECT MAX(TIME_OF_SETTING) FROM TOP_STORIES") // check for the newest timestamp in the db

			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				err = rows.Scan(&lastDBOutTime)
				if err != nil {
					log.Println(err)
				}
			}
		}

		if time.Now().UnixMilli()-lastDBOutTime < time.Hour.Milliseconds() { // taking data from the db because it hasn't been an hour
			var topStories Stories
			rows, err := db.Query("SELECT TITLE, SCORE FROM TOP_STORIES")
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				cachedStory := Story{}
				err = rows.Scan(&cachedStory.Title, &cachedStory.Score)
				if err != nil {
					log.Println(err)
				}
				topStories.AllStories = append(topStories.AllStories, cachedStory)
			}
			storiesJSON, err := json.Marshal(topStories)
			if err != nil {
				log.Println(err)
			}
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, storiesJSON, "", "    "); err != nil {
				log.Println(err)
			}
			data := StoriesPageData{PageTitle: "Top stories of Hacker News FROM DB", Links: topStories.AllStories}
			tmpl := template.Must(template.ParseFiles("templ.html"))
			tmpl.Execute(w, data)

		} else {

			// if it has been an hour since last pull from the HN api

			log.Println(r.Method)
			if r.Method != http.MethodGet {
				http.Error(w, "Only methods are allowed", http.StatusBadRequest)
				return
			}
			log.Println("HandleTop called")
			res, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")

			if err != nil {
				log.Println(err)
			}
			var payload IDPayload
			json.NewDecoder(res.Body).Decode(&payload.IDs)
			var URLs [10]string
			for i := 0; i < 10; i++ {
				URLs[i] = "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(payload.IDs[i]) + ".json?print=pretty"
			}
			var topStories Stories
			for i := 0; i < 10; i++ {
				res, err := http.Get(URLs[i])
				if err != nil {
					log.Println(err)
				}
				topStories.AllStories = append(topStories.AllStories, Story{})
				json.NewDecoder(res.Body).Decode(&topStories.AllStories[i])
			}

			storiesJSON, err := json.Marshal(topStories)
			if err != nil {
				log.Println(err)
			}
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, storiesJSON, "", "    "); err != nil {
				log.Println(err)
			}
			data := StoriesPageData{PageTitle: "Top stories of Hacker News", Links: topStories.AllStories}
			tmpl := template.Must(template.ParseFiles("templ.html"))
			tmpl.Execute(w, data)

			rows, err := db.Query("SELECT COUNT(id) FROM TOP_STORIES")
			if err != nil {
				log.Println(err)
			}
			var count int
			for rows.Next() {
				err = rows.Scan(&count)
				if err != nil {
					log.Print(err)
				}
			}
			if err != nil {
				log.Print(err)
			}
			if count > 0 {

				_, errDB := db.Exec("DELETE FROM TOP_STORIES")
				if errDB != nil {
					log.Println(err)
				}

			}
			for _, story := range topStories.AllStories {
				timeInMillis := time.Now().UnixMilli()
				_, errDB := db.Exec("INSERT INTO TOP_STORIES(TITLE, SCORE, TIME_OF_SETTING) values(?,?,?)", story.Title, story.Score, timeInMillis)
				if errDB != nil {
					log.Println(err)
				}
			}
		}
	}

}
