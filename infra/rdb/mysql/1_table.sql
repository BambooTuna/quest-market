DROP SCHEMA IF EXISTS market;
CREATE SCHEMA market;
USE market;

CREATE TABLE `account_credentials` (
    `account_id` VARCHAR(255) NOT NULL,
    `mail` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`account_id`),
    UNIQUE KEY (`mail`)
);

CREATE TABLE `linked_user_credentials` (
    `id` VARCHAR(255) NOT NULL,
    `service_id` VARCHAR(255) NOT NULL,
    `service_name` VARCHAR(255) NOT NULL,
    `mail` VARCHAR(255),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`service_id`, `service_name`),
    UNIQUE KEY (`mail`)
);

CREATE TABLE `product_display` (
    `id` VARCHAR(255) NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `detail` VARCHAR(255) NOT NULL,
    `price` bigint NOT NULL,
    `presenter_id` VARCHAR(255) NOT NULL,
    `state` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);
