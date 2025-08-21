-- name: CreateFeedFollow :many

WITH insert_feed_follows AS (
  INSERT INTO feed_follows (
    id, created_at, updated_at, user_id, feed_id
  ) VALUES ( $1, $2, $3, $4, $5 )
    RETURNING *
)
SELECT insert_feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM insert_feed_follows
INNER JOIN users ON insert_feed_follows.user_id = users.id
INNER JOIN feeds ON insert_feed_follows.feed_id = feeds.id;
