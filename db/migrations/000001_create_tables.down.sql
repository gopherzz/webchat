CREATE TABLE `messages` (
  `id` uuid PRIMARY KEY NOT NULL,
  `sender_id` uuid NOT NULL,
  `data` text NOT NULL
);

CREATE TABLE `users` (
  `id` uuid PRIMARY KEY NOT NULL,
  `name` varchar(50) NOT NULL
);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `messages` (`sender_id`);
