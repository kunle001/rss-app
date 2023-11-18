package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kunle001/rss-app/internal/database"
)

func startScrapping(
	db *database.Queries, 
	concurrency int, 
	timeBetweeenRequest time.Duration,
){
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweeenRequest)
	ticker := time.NewTicker(timeBetweeenRequest);

	for ; ; <-ticker.C{//Ensures the job starts instantly 
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(), 
			int32(concurrency),
		)

		if err !=nil{
			log.Println("error fetching feeds", err)
			continue
		};

		wg := &sync.WaitGroup{};

		for _, feed := range feeds{
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}

		wg.Wait()//Blocking and waiting for go routine to get done
	}

}

func scrapeFeed( db *database.Queries,wg *sync.WaitGroup, feed database.Feed){
	defer wg.Done()

	_, err :=db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err !=nil{
		log.Println("Error making feed as fetched", err)
		return
	}

	rssFeed, err :=urlToFeed(feed.Url)

	if err!= nil{
		log.Println("Error converting url to feed", err)
		return 
	}

	for _, item := range rssFeed.Channel.Item{
		descrpition := sql.NullString{};

		if item.Description != ""{
			descrpition.String = item.Description
			descrpition.Valid = true
		}
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err !=nil{
			log.Println("Error occured while parsing the time", err)
		}

		_, err =db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(), 
			UpdatedAt: time.Now(), 
			Url: item.Link, 
			Title: item.Title,
			Description: descrpition,
			PublishedAt: pubAt, 
			FeedID: feed.ID,
		});

		if err!=nil{
			if strings.Contains(err.Error(), "duplicate key"){
				continue
			}
			log.Printf("Error created feed %v Because %v", feed, err )
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}