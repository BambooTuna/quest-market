DROP SCHEMA IF EXISTS market;
CREATE SCHEMA market;
USE market;

CREATE TABLE `account_credentials` (
    `account_id` VARCHAR(255) NOT NULL,
    `mail` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`account_id`),
    UNIQUE KEY (`mail`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `product_details` (
  `product_id` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `detail` varchar(255) NOT NULL,
  `price` bigint(20) NOT NULL,
  `presenter_id` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
