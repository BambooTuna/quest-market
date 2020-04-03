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
-- INSERT INTO `account_credentials` (`account_id`, `mail`, `password`) VALUES ("f0c28384-3aa4-3f87-9fba-66a0aa62c504", "bambootuna@gmail.com", "d74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1");

CREATE TABLE `item_details` (
  `item_id` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `detail` varchar(255) NOT NULL,
  `price` bigint(20) NOT NULL,
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- INSERT INTO `item_details` (`item_id`, `title`, `detail`, `price`) VALUES ("1", "title", "detail", 100);

CREATE TABLE `contract_details` (
  `item_id` varchar(255) NOT NULL,
  `purchaser_account_id` varchar(255),
  `seller_account_id` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `created_at` timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  `updated_at` timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- INSERT INTO `contract_details` (`item_id`, `purchaser_account_id`, `seller_account_id`, `state`) VALUES ("1", "purchaser_account_id", "seller_account_id", "open");

/*
transaction_type: deposit, withdraw, canceled_deposit, canceled_withdraw
currency: jpy, usd
*/
CREATE TABLE `money_transaction` (
  `transaction_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `transaction_type` varchar(255) NOT NULL,
  `account_id` varchar(255) NOT NULL,
  `currency` varchar(255) NOT NULL,
  `real_part` bigint(20) NOT NULL,
  `exponent_part` bigint(20) NOT NULL,
  `created_at` timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  PRIMARY KEY (`transaction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*
INSERT INTO `money_transaction` (`transaction_type`, `account_id`, `currency`, `real_part`, `exponent_part`) VALUES
("deposit", "f0c28384-3aa4-3f87-9fba-66a0aa62c504", "jpy", 10000, 1);
*/
