-- name: FindItem :one
SELECT * FROM VMT_Items WHERE ID = ?;

-- name: CreateItem :exec
INSERT INTO VMT_Items (ID, Title, Description, IsGood, CreatedAt) VALUES (?,?,?,?,?);