package main

import (
	"database/sql"
	"time"

	"github.com/codezera11/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
}

type UsersFeed struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUsersFeedToUsersFeed(usersFeed database.UsersFeed) UsersFeed {
	return UsersFeed{
		ID:        usersFeed.ID,
		FeedID:    usersFeed.FeedID,
		UserID:    usersFeed.UserID,
		CreatedAt: usersFeed.CreatedAt,
		UpdatedAt: usersFeed.UpdatedAt,
	}
}

func databaseUsersFeedsToUsersFeeds(dbUsersFeeds []database.UsersFeed) []UsersFeed {
	usersFeeds := make([]UsersFeed, len(dbUsersFeeds))

	for _, dbUserFeed := range dbUsersFeeds {
		usersFeeds = append(usersFeeds, databaseUsersFeedToUsersFeed(dbUserFeed))
	}

	return usersFeeds
}

type Post struct {
	ID          uuid.UUID      `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Title       string         `json:"title"`
	Url         sql.NullString `json:"url"`
	Description sql.NullString `json:"description"`
	PublishedAt sql.NullTime   `json:"published_at"`
	FeedID      uuid.UUID      `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: post.Description,
		PublishedAt: post.PublishedAt,
		FeedID:      post.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := make([]Post, len(dbPosts))

	for _, dbPost := range dbPosts {
		posts = append(posts, Post(dbPost))
	}

	return posts
}
