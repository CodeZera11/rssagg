-- +goose Up
CREATE TABLE users_feeds (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY(feed_id) REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE users_feeds;