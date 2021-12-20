-- name: insert-data
INSERT INTO scraper.products (title, image_url, price, description, url) VALUES (?,?,?,?,?);

-- name: get-data
SELECT
    title Title,
    image_url ImageURL,
    price Price,
    description Description,
    url URL,
    insertion_date InsertionDate
FROM scraper.products
WHERE url = ?;
