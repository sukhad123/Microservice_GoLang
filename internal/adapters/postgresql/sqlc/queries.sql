-- name: ListProducts :many
Select 
* From products;
 
-- name: CreateProduct :one
INSERT INTO products (name, price, quantity, created_at)
VALUES ($1, $2, $3, NOW())
RETURNING *;
