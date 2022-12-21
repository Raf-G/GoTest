-- Struct table `users`
CREATE TABLE `users` (
    `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT,
    `login` char(255) COLLATE NOT NULL,
    `surname` char(255) NOT NULL,
    `name` char(255) NOT NULL,
    `role` tinyint(5) UNSIGNED DEFAULT NULL
);

ALTER TABLE `users`
    ADD PRIMARY KEY (`id`),
    ADD KEY `role` (`role`);

ALTER TABLE `users`
    ADD CONSTRAINT `users_ibfk_1` FOREIGN KEY (`role`) REFERENCES `roles` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;
    COMMIT;

-- Struct table `roles`
CREATE TABLE `roles` (
    `id` tinyint(5) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` char(255) COLLATE NOT NULL
);

ALTER TABLE `roles`
    ADD PRIMARY KEY (`id`),
    ADD UNIQUE KEY `name` (`name`),
    ADD KEY `id` (`id`);

-- Struct table `products`
CREATE TABLE `products` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` char(255) COLLATE NOT NULL,
    `description` text COLLATE DEFAULT NULL,
    `price` double DEFAULT NULL
);

ALTER TABLE `products`
    ADD PRIMARY KEY (`id`);

-- Struct table `reviews`
CREATE TABLE `reviews` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(10) UNSIGNED NOT NULL,
    `product_id` int(10) UNSIGNED NOT NULL,
    `description` text COLLATE DEFAULT NULL,
    `grade` smallint(6) DEFAULT NULL
);

ALTER TABLE `reviews`
    ADD PRIMARY KEY (`id`);

-- Struct table `products_basket`
CREATE TABLE `products_basket` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `basket_id` int(10) UNSIGNED DEFAULT NULL,
    `product_id` int(10) UNSIGNED DEFAULT NULL,
    `count` int(10) UNSIGNED NOT NULL
);

ALTER TABLE `products_basket`
    ADD PRIMARY KEY (`id`),
    ADD KEY `basket_id` (`basket_id`),
    ADD KEY `product_id` (`product_id`);

-- Struct table `baskets`
CREATE TABLE `baskets` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(10) UNSIGNED NOT NULL
);

ALTER TABLE `baskets`
    ADD PRIMARY KEY (`id`);

-- Struct table `products_order`
CREATE TABLE `products_order` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `order_id` int(10) UNSIGNED DEFAULT NULL,
    `product_id` int(10) UNSIGNED DEFAULT NULL,
    `count` int(10) UNSIGNED NOT NULL,
    `price` int(10) UNSIGNED NOT NULL
);

ALTER TABLE `products_order`
    ADD PRIMARY KEY (`id`),
    ADD KEY `order_id` (`order_id`),
    ADD KEY `product_id` (`product_id`);

-- Struct table `orders`
CREATE TABLE `orders` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(10) UNSIGNED NOT NULL,
    `status` int(10) UNSIGNED DEFAULT NULL
);

ALTER TABLE `orders`
    ADD PRIMARY KEY (`id`),
    ADD KEY `status` (`status`),
    ADD KEY `id` (`id`);

-- Struct table `status`
CREATE TABLE `status` (
    `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` char(255) COLLATE DEFAULT NULL
);

ALTER TABLE `status`
    ADD PRIMARY KEY (`id`),
    ADD KEY `id` (`id`);

-- --------------------------------------------------------------------------------
ALTER TABLE `products_basket`
    ADD CONSTRAINT `products_basket_ibfk_1` FOREIGN KEY (`basket_id`) REFERENCES `baskets` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `products_basket_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE `products_order`
    ADD CONSTRAINT `products_order_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `products_order_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE `orders`
    ADD CONSTRAINT `order_ibfk_1` FOREIGN KEY (`status`) REFERENCES `status` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

-- --------------------------------------------------------------------------------
-- Default data 

INSERT INTO `users` (`id`, `surname`, `name`, `role`) VALUES
    (1, 'usergasf', 'brhtrh', 1),
    (2, 'ivan', 'totot', 1),
    (3, 'ivan2', 'totot2', 2),
    (4, 'ivan3', 'totot3', 3);

INSERT INTO `roles` (`id`, `name`) VALUES
    (1, 'administrator'),
    (2, 'user');

INSERT INTO `products` (`id`, `name`, `description`, `price`) VALUES
    (1, 'car', 'car is a cool', 100);

INSERT INTO `reviews` (`id`, `user_id`, `product_id`, `description`, `grade`) VALUES
    (1, 1, 1, 'wqwqwq', 3),
    (2, 1, 1, 'wqwqwq', 1),
    (3, 1, 1, 'wqwqwq', 5);