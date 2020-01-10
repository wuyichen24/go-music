-- Dumping database structure for gomusic
CREATE DATABASE IF NOT EXISTS `gomusic` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */;
USE `gomusic`;

-- Dumping structure for table gomusic.customers
CREATE TABLE IF NOT EXISTS `customers` (
  `id`            int(11)      NOT NULL AUTO_INCREMENT,
  `firstname`     varchar(50)  NOT NULL DEFAULT '0',
  `lastname`      varchar(50)  NOT NULL DEFAULT '0',
  `email`         varchar(100) NOT NULL DEFAULT '0',
  `pass`          varchar(100) NOT NULL DEFAULT '0',
  `cc_customerid` varchar(50)  NOT NULL DEFAULT '0',
  `loggedin`      tinyint(1)   NOT NULL DEFAULT '0',
  `created_at`    timestamp    NULL     DEFAULT CURRENT_TIMESTAMP,
  `updated_at`    timestamp    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`    timestamp    NULL     DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;

-- Dumping structure for table gomusic.orders
CREATE TABLE IF NOT EXISTS `orders` (
  `id`            int(11)   NOT NULL AUTO_INCREMENT,
  `customer_id`   int(11)   NOT NULL,
  `product_id`    int(11)   NOT NULL,
  `price`         int(11)   NOT NULL,
  `purchase_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at`    timestamp NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at`    timestamp NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`    timestamp NULL     DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- Dumping structure for table gomusic.products
CREATE TABLE IF NOT EXISTS `products` (
  `id`          int(11) NOT NULL AUTO_INCREMENT,
  `image`       varchar(100)   DEFAULT NULL,
  `smallimg`    varchar(100)   DEFAULT NULL,
  `imgalt`      varchar(50)    DEFAULT NULL,
  `description` text,
  `productname` varchar(50)    DEFAULT NULL,
  `price`       float          DEFAULT NULL,
  `promotion`   float          DEFAULT NULL,
  `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at`  timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;
