CREATE TABLE IF NOT EXISTS motting.users (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE motting.phrases (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` varchar(24) DEFAULT NULL,
  `text` varchar(128) DEFAULT NULL,
  `author` varchar(24) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_phrases_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE motting.push_times (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` varchar(24) DEFAULT NULL,
  `push_at` varchar(5) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_push_times_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE webpush.subscriptions (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` varchar(24) DEFAULT NULL,
  `endpoint` varchar(2048) DEFAULT NULL,
  `p256dh` varchar(255) DEFAULT NULL,
  `auth` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_subscriptions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;