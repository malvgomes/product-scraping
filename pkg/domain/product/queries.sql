-- name: insert-data
INSERT INTO scraper.products (title, image_url, price, description, url) VALUES (?,?,?,?,?)
ON DUPLICATE KEY UPDATE title = VALUES(title), image_url = VALUES(image_url), price = VALUES(price), description = VALUES(description), date = NOW();

-- name: get-data
SELECT
    title Title,
    image_url ImageURL,
    price Price,
    description Description,
    url URL,
    date InsertionDate
FROM scraper.products
WHERE url = ?;
