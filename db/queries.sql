-- name: InsertGame :one
INSERT INTO games (game_title) VALUES (sqlc.arg(gameTitle)) RETURNING id;
