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
  `created_at` timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  `updated_at` timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
transaction_type: waiting_for_payment, waiting_to_receive, complete
*/
CREATE TABLE `product_transaction` (
  `transaction_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `transaction_type` varchar(255) NOT NULL,
  `product_id` varchar(255) NOT NULL,
  `sender_account_id` varchar(255) NOT NULL,
  `receiver_account_id` varchar(255) NOT NULL,
  `created_at` timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  PRIMARY KEY (`transaction_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
