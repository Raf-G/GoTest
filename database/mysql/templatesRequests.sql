-- Get all carts with the total amount of their items
SELECT *, price * count AS total_price, SUM(price * count) AS busket_price
FROM baskets
         JOIN products_baskets ON baskets.id = products_baskets.basket_id
         JOIN products ON products_baskets.product_id = products.id
GROUP BY products_baskets.basket_id

-- Get all users who do not have items in their cart
SELECT *
FROM users
         JOIN baskets ON users.id = baskets.user_id
         LEFT JOIN products_baskets ON baskets.id = products_baskets.basket_id
WHERE products_baskets.id IS NULL


