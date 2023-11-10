package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/kunle001/rss-app/internal/database"
)

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Name string	`json:"name"`
	ApiKey string	`json:"ApiKey"`
}

type Feed struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Name string	`json:"name"`
	Url string	`json:"URL"`
	UserID uuid.UUID `json:"UserID"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
	CreatedAt: dbUser.CreatedAt,
	UpdatedAt: dbUser.UpdatedAt,
	Name: dbUser.Name,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
	CreatedAt: dbFeed.CreatedAt,
	UpdatedAt: dbFeed.UpdatedAt,
	Name: dbFeed.Name,
	Url: dbFeed.Url,
	UserID: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed:= range dbFeeds{
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	};

	return feeds
}