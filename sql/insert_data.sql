INSERT INTO `customers` (`id`, `firstname`, `lastname`, `email`, `pass`, `cc_customerid`, `loggedin`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'Mal', 'Zein', 'mal.zein@email.com', '$2a$10$ZeZI4pPPlQg89zfOOyQmiuKW9Z7pO9/KvG7OfdgjPAZF0Vz9D8fhC', 'cus_EL08toK8pfDcom', 0, '2018-08-14 07:52:54', '2019-03-03 19:01:16', NULL),
	(2, 'River', 'Sam', 'river.sam@email.com', '$2a$10$mNbCLmfCAc0.4crDg3V3fe0iO1yr03aRfE7Rr3vdfKMGVnnzovCZq', '', 0, '2018-08-14 07:52:55', '2019-01-12 22:39:01', NULL),
	(3, 'Jayne', 'Ra', 'jayne.ra@email.com', '$2a$10$ZeZI4pPPlQg89zfOOyQmiuKW9Z7pO9/KvG7OfdgjPAZF0Vz9D8fhC', 'cus_EL4GpQmVjwvUUZ', 0, '2018-08-14 07:52:55', '2019-01-13 21:56:05', NULL),
	(19, 'John', 'Doe', 'john.doe@bla.com', '$2a$10$T4c8rmpbgKrUA0sIqtHCaO0g2XGWWxFY4IGWkkpVQOD/iuBrwKrZu', '', 0, '2019-01-13 08:43:44', '2019-01-13 15:12:25', NULL);

INSERT INTO `orders` (`id`, `customer_id`, `product_id`, `price`, `purchase_date`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 1, 90, '2018-12-29 23:34:32', '2018-12-29 23:35:36', '2018-12-29 23:35:36', NULL),
	(2, 1, 2, 299, '2018-12-29 23:34:53', '2018-12-29 23:35:48', '2018-12-29 23:35:48', NULL),
	(3, 1, 3, 16000, '2018-12-29 23:35:05', '2018-12-29 23:35:57', '2018-12-29 23:35:57', NULL),
	(4, 2, 1, 95, '2018-12-29 23:36:18', '2018-12-29 23:36:18', '2018-12-29 23:36:18', NULL),
	(5, 2, 2, 299, '2018-12-29 23:36:39', '2018-12-29 23:36:39', '2018-12-29 23:36:39', NULL),
	(6, 2, 4, 205, '2018-12-29 23:37:01', '2018-12-29 23:38:13', '2018-12-29 23:38:13', NULL),
	(7, 3, 4, 210, '2018-12-29 23:37:28', '2018-12-29 23:38:19', '2018-12-29 23:38:19', NULL),
	(8, 3, 5, 200, '2018-12-29 23:37:41', '2018-12-29 23:38:28', '2018-12-29 23:38:28', NULL),
	(9, 3, 6, 1000, '2018-12-29 23:37:54', '2018-12-29 23:38:32', '2018-12-29 23:38:32', NULL),
	(10, 19, 6, 1000, '2018-12-29 23:37:54', '2019-01-13 00:44:55', '2019-01-13 00:44:55', NULL),
	(11, 1, 3, 17000, '2018-12-29 23:37:56', '2019-01-14 06:03:08', '2019-01-14 06:03:08', NULL);

INSERT INTO `products` (`id`, `image`, `smallimg`, `imgalt`, `description`, `productname`, `price`, `promotion`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'img/strings.png', 'img/img-small/strings.png', 'string', '', 'Strings', 100, NULL, '2018-08-14 07:54:19', '2019-01-11 00:28:40', NULL),
	(2, 'img/redguitar.jpeg', 'img/img-small/redguitar.jpeg', 'redg', '', 'Red Guitar', 299, 240, '2018-08-14 07:54:20', '2019-01-11 00:29:11', NULL),
	(3, 'img/drums.jpg', 'img/img-small/drums.jpg', 'drums', '', 'Drums', 17000, NULL, '2018-08-14 07:54:20', '2019-01-11 22:05:42', NULL),
	(4, 'img/flute.jpeg', 'img/img-small/flute.jpeg', 'flute', '', 'Flute', 210, 190, '2018-08-14 07:54:20', '2019-01-11 00:29:53', NULL),
	(5, 'img/blackguitar.jpeg', 'img/img-small/blackguitar.jpeg', 'Black guitar', '', 'Black Guitar', 200, NULL, '2018-08-14 07:54:20', '2019-01-11 00:30:12', NULL),
	(6, 'img/saxophone.jpeg', 'img/img-small/saxophone.jpeg', 'Saxophone', '', 'Saxophone', 1000, 980, '2018-08-14 07:54:20', '2019-01-11 00:30:35', NULL);
