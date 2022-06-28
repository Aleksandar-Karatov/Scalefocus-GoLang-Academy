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
	"week9Lecture26Task/week9/lecture26/dataStructs"
	"week9Lecture26Task/week9/lecture26/repository"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "DBL26.db")
	if err != nil {
		log.Fatal(err)
	}
	var dbc repository.DBcontext
	dbc.DB = db
	mux := http.NewServeMux()

	defer db.Close()
	mux.Handle("/api/top", HandleTop(dbc, "https://hacker-news.firebaseio.com/v0/"))
	log.Fatal(http.ListenAndServe(":9000", mux))
}

func HandleTop(dbc repository.DBcontext, dest string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		lastDBOutTime, err := dbc.GetLatestTimestamp()
		if err != nil {
			log.Println(err)
		}

		if time.Now().UnixMilli()-lastDBOutTime < time.Hour.Milliseconds() { // taking data from the db because it hasn't been an hour
			log.Println(r.Method)
			storiesJSON, err, topStories := dbc.GetAllStoriesAsJSON()
			if err != nil {
				log.Println(err)
			}
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, storiesJSON, "", "    "); err != nil {
				log.Println(err)
			}
			data := dataStructs.StoriesPageData{PageTitle: "Top stories of Hacker News FROM DB", Links: topStories.AllStories}
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
			var res *http.Response
			var err error
			if dest == "https://hacker-news.firebaseio.com/v0/" {
				res, err = http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
			} else {
				res, err = http.Get(dest + "/IDs")
			}

			if err != nil {
				log.Println(err)
			}

			var payload dataStructs.IDPayload
			json.NewDecoder(res.Body).Decode(&payload.IDs)
			var URLs [10]string
			for i := 0; i < 10; i++ {
				if dest == "https://hacker-news.firebaseio.com/v0/" {
					URLs[i] = "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(payload.IDs[i]) + ".json?print=pretty"
				} else {
					URLs[i] = dest + "/" + strconv.Itoa(payload.IDs[i])
				}
			}
			var topStories dataStructs.Stories
			for i := 0; i < 10; i++ {
				res, err := http.Get(URLs[i])
				if err != nil {
					log.Println(err)
				}
				topStories.AllStories = append(topStories.AllStories, dataStructs.Story{})
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
			data := dataStructs.StoriesPageData{PageTitle: "Top stories of Hacker News", Links: topStories.AllStories}
			tmpl := template.Must(template.ParseFiles("templ.html"))
			tmpl.Execute(w, data)
			errDB := dbc.AddIntoDB(topStories)
			if errDB != nil {
				log.Println(errDB)
			}

		}
	}

}
