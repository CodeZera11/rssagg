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

-- name: DeleteFollowFeed :one
DELETE FROM users_feeds
WHERE id = $1
RETURNING *;

-- name: GetFollowedFeeds :many
SELECT * FROM users_feeds
WHERE user_id = $1;