-- name: CreateVibePointTransaction :one
INSERT INTO vibe_point_transactions (
  awarded_user, points_awarded, reason, created, modified_at
) VALUES ( 
  $1, $2, $3, NOW(), NOW()
)
RETURNING *;
