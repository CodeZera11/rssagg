-- name: CreateFeed :one
INSERT INTO feeds(id, name, url, user_id, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: FollowFeed :one
INSERT INTO users_feeds(id, feed_id, user_id, created_at, updated_at)
VALUES($1, $2, $3, $4, $5)
RETURNING *;