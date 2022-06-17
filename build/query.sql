-- name: ListBooks :many
SELECT *
FROM book;
-- name: ListReviews :many
SELECT *
FROM review;
-- name: ListAuthors :many
SELECT *
FROM author;