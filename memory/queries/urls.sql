-- name: CreateUrl :exec
INSERT INTO urls (shortened_code, original_url)
VALUES ($1, $2);


-- name: GetUrl :one
SELECT original_url FROM urls
WHERE shortened_code = $1;
