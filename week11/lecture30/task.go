package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
	"week11Lecture30Task/protoserv"
	"week11Lecture30Task/repositoryWithCodeGenFinalHW"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "L30DB.db")
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	defer db.Close()
	mux.Handle("/api/top", HandleTop(db, "https://hacker-news.firebaseio.com/v0/"))
	log.Fatal(http.ListenAndServe(":9000", mux))
}

type NewsID = int

type IDPayload struct {
	IDs []NewsID
}

func HandleTop(db *sql.DB, dest string) http.HandlerFunc {

	ctx := context.Background()
	queries := repositoryWithCodeGenFinalHW.New(db)
	return func(w http.ResponseWriter, r *http.Request) {
		lastDBOutTime, err := queries.GetLatestTime(ctx)
		if err != nil {
			log.Println(err)
		}

		log.Println(time.Now().UnixMilli() - lastDBOutTime)
		if time.Now().UnixMilli()-lastDBOutTime < time.Hour.Milliseconds() { // taking data from the db because it hasn't been an hour
			log.Println(r.Method)
			log.Println("Pulling from db")
			storiesData, err := queries.GetDataForStories(ctx)
			if err != nil {
				log.Println(err)
			}
			for _, item := range storiesData.AllStoriesInRows {
				SendMessage(item)

			}
			storiesJSON, err := json.Marshal(storiesData)

			if err != nil {
				log.Println(err)
			}
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, storiesJSON, "", "    "); err != nil {
				log.Println(err)
			}
			data := repositoryWithCodeGenFinalHW.StoriesPageDataFromDB{PageTitle: "Top stories of Hacker News FROM DB", Links: storiesData.AllStoriesInRows}
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

			var payload IDPayload
			json.NewDecoder(res.Body).Decode(&payload.IDs)
			var URLs [10]string
			for i := 0; i < 10; i++ {
				if dest == "https://hacker-news.firebaseio.com/v0/" {
					URLs[i] = "https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(payload.IDs[i]) + ".json?print=pretty"
				} else {
					URLs[i] = dest + "/" + strconv.Itoa(payload.IDs[i])
				}
			}
			var topStories repositoryWithCodeGenFinalHW.DataPayloadForRows
			for i := 0; i < 10; i++ {
				res, err := http.Get(URLs[i])
				if err != nil {
					log.Println(err)
				}
				topStories.AllStoriesInRows = append(topStories.AllStoriesInRows, repositoryWithCodeGenFinalHW.GetDataForStoriesRow{})
				json.NewDecoder(res.Body).Decode(&topStories.AllStoriesInRows[i])
			}
			for _, item := range topStories.AllStoriesInRows {
				SendMessage(item)

			}

			storiesJSON, err := json.Marshal(topStories)
			if err != nil {
				log.Println(err)
			}
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, storiesJSON, "", "    "); err != nil {
				log.Println(err)
			}
			data := repositoryWithCodeGenFinalHW.StoriesPageDataFromDB{PageTitle: "Top stories of Hacker News FROM WEB", Links: topStories.AllStoriesInRows}
			tmpl := template.Must(template.ParseFiles("templ.html"))
			tmpl.Execute(w, data)

			errDB := queries.DeleteEverything(ctx)
			if err != nil {
				log.Println(errDB)
			}
			for _, story := range topStories.AllStoriesInRows {
				_, err := queries.InsertIntoDB(ctx, repositoryWithCodeGenFinalHW.InsertIntoDBParams{Title: story.Title, Score: story.Score, TimeOfSetting: time.Now().UnixMilli()})
				if err != nil {
					log.Println(err)
				}
			}

		}
	}

}

var Addr = flag.String("addr", "localhost:50051", "the address to connect to")

func SendMessage(story repositoryWithCodeGenFinalHW.GetDataForStoriesRow) {

	var Conn, Err = grpc.Dial(*Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if Err != nil {
		log.Println(Err)
	}
	var C = protoserv.NewDataGetterClient(Conn)
	var Ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var title = flag.String("title"+story.Title+strconv.Itoa(rand.Intn(math.MaxInt)), story.Title, "title of the story")
	var score = flag.Int64("score"+strconv.Itoa(rand.Intn(math.MaxInt)), int64(story.Score), "score of the story")
	flag.Parse()

	r, err := C.ShowData(Ctx, &protoserv.DataRequest{Title: *title, Score: *score})
	if err != nil {
		log.Println(err)
	}
	log.Println("Title is: "+r.GetTitle(), "Score is: ", r.GetScore())

	defer Conn.Close()

}
