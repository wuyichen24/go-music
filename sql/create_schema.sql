CREATE TABLE `account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstname`     varchar(50)       DEFAULT NULL,
  `lastname`      varchar(50)       DEFAULT NULL,
  `email`         varchar(100)      DEFAULT NULL,
  `pass`          varchar(100)      DEFAULT NULL,
  `cc_customerid` varchar(50)       DEFAULT NULL,
  `loggedin`      tinyint(1)        DEFAULT NULL,
  `created_at`    timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`    timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at`    timestamp    NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id`    int(11),
  `product_id`     int(11),
  `price`       int(11),
  `purchased_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at`     timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`     timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at`     timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

CREATE TABLE `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `image`       varchar(255)  NULL DEFAULT NULL,
  `imgalt`      varchar(50)   NULL DEFAULT NULL,
  `description` text          ,
  `productname` varchar(100)  NULL DEFAULT NULL,
  `price`       float DEFAULT NULL,
  `promotion`   float DEFAULT NULL,
  `created_at`  timestamp     NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  timestamp     NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at`  timestamp     NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
