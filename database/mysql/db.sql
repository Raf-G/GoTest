-- Delete tables
DROP TABLE IF EXISTS products_baskets;
DROP TABLE IF EXISTS baskets;
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS products_orders;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS statuses;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;

-- Struct table `users`
CREATE TABLE `users`
(
    `id`       int UNSIGNED     NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `login`    VARCHAR(50)      NOT NULL,
    `surname`  VARCHAR(100)     NOT NULL,
    `name`     VARCHAR(100)     NOT NULL,
    `password` VARCHAR(50)      NOT NULL,
    `role_id`  tinyint UNSIGNED NOT NULL
);

-- Struct table `roles`
CREATE TABLE `roles`
(
    `id`   tinyint UNSIGNED NOT NULL PRIMARY KEY,
    `name` VARCHAR(50)      NOT NULL
);

-- Struct table `reviews`
CREATE TABLE `reviews`
(
    `id`          int UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id`     int UNSIGNED NOT NULL,
    `product_id`  int UNSIGNED NOT NULL,
    `description` text DEFAULT NULL,
    `grade`       smallint     NOT NULL
);

-- Struct table `products_baskets`
CREATE TABLE `products_baskets`
(
    `id`         int UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `basket_id`  int UNSIGNED NOT NULL,
    `product_id` int UNSIGNED NOT NULL,
    `count`      int UNSIGNED NOT NULL
);

-- Struct table `baskets`
CREATE TABLE `baskets`
(
    `id`      int UNSIGNED NOT NULL PRIMARY KEY,
    `user_id` int UNSIGNED NOT NULL
);

-- Struct table `products_orders`
CREATE TABLE `products_orders`
(
    `id`         int UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `order_id`   int UNSIGNED NOT NULL,
    `product_id` int UNSIGNED NOT NULL,
    `count`      int UNSIGNED NOT NULL,
    `price`      int UNSIGNED NOT NULL
);

-- Struct table `orders`
CREATE TABLE `orders`
(
    `id`        int UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id`   int UNSIGNED NOT NULL,
    `status_id` int UNSIGNED NOT NULL DEFAULT 1
);

-- Struct table `status`
CREATE TABLE statuses
(
    `id`   int UNSIGNED NOT NULL PRIMARY KEY,
    `name` VARCHAR(50)  NOT NULL
);

-- Struct table `products`
CREATE TABLE `products`
(
    `id`          int UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name`        VARCHAR(100) NOT NULL,
    `description` text DEFAULT NULL,
    `price`       int  DEFAULT NULL
);

-- --------------------------------------------------------------------------------
ALTER TABLE `users`
    ADD CONSTRAINT `fk_users_role_id_roles_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `reviews`
    ADD CONSTRAINT `fk_reviews_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    ADD CONSTRAINT `fk_reviews_product_id_products_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `products_baskets`
    ADD CONSTRAINT `products_baskets_basket_id_baskets_id` FOREIGN KEY (`basket_id`) REFERENCES `baskets` (`id`),
    ADD CONSTRAINT `products_baskets_product_id_products_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `baskets`
    ADD CONSTRAINT `baskets_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `products_orders`
    ADD CONSTRAINT `products_orders_order_id_orders_id` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
    ADD CONSTRAINT `products_orders_product_id_products_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `orders`
    ADD CONSTRAINT `orders_user_id_users_id` FOREIGN KEY (`user_id`) REFERENCES users (`id`),
    ADD CONSTRAINT `orders_status_id_statuses_id` FOREIGN KEY (`status_id`) REFERENCES statuses (`id`);

-- --------------------------------------------------------------------------------
-- Default data
INSERT INTO `roles` (`id`, `name`)
VALUES (1, 'administrator'),
       (2, 'user');

INSERT INTO `users` (`id`, `login`, `surname`, `name`, `role_id`, `password`)
VALUES (1, 'admin', 'Nikita', 'Tranche', 1, '12345678'),
       (2, 'tes2', 'ivan', 'ivanovhich', 1, 'qwerty'),
       (3, 'tes3', 'ivan2', 'kanovich', 2, 'asdf');

INSERT INTO `products` (`id`, `name`, `description`, `price`)
VALUES (1, 'car', 'car is a cool', 100),
       (2, 'train', 'small train', 150),
       (3, 'plane', 'aero-bus', 200);

INSERT INTO `reviews` (`id`, `user_id`, `product_id`, `description`, `grade`)
VALUES (1, 1, 1, 'wqwqwq', 3),
       (2, 1, 1, 'wqwqwq', 1),
       (3, 1, 1, 'wqwqwq', 5);

INSERT INTO `baskets` (`id`, `user_id`)
VALUES (1, 1),
       (2, 2),
       (3, 3);

INSERT INTO `products_baskets` (`basket_id`, `product_id`, `count`)
VALUES (1, 1, 2),
       (1, 2, 4),
       (1, 3, 5),
       (2, 2, 4);

INSERT INTO `statuses` (`id`, `name`)
VALUES (1, 'В обработке'),
       (2, 'Сборка'),
       (3, 'Отправлен'),
       (4, 'Доставлен'),
       (5, 'Отменен');

INSERT INTO `orders` (`id`, `user_id`, `status_id`)
VALUES (1, 1, 1),
       (2, 2, 2),
       (3, 3, 3),
       (4, 2, 1);

INSERT INTO `products_orders` (`id`, `order_id`, `product_id`, `count`, `price`)
VALUES (1, 1, 1, 3, 3000),
       (2, 2, 2, 3, 111),
       (3, 3, 3, 4, 333),
       (4, 2, 1, 6, 200);

