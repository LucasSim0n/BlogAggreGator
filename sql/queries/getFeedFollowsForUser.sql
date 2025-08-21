-- name: GetFeedFollowsForUser :many

WITH ff AS (
  SELECT feed_id FROM feed_follows WHERE user_id = $1
)
SELECT ff.feed_id, feeds.name AS feed_name, users.name AS user_name
FROM ff
INNER JOIN feeds ON feeds.id = ff.feed_id
INNER JOIN users ON users.id = $1;

