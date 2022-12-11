CREATE TABLE users (
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) UNIQUE NOT NULL, -- `nickname`
    `email` VARCHAR(255) UNIQUE NOT NULL, -- `email address`
    `password` VARCHAR(255) NOT NULL, -- `password hash`
    PRIMARY KEY (`id`)
);


CREATE TABLE articles (
    `id` INT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(100) NOT NULL, -- `title of article`
	`announce` VARCHAR(255), -- `anouncement of article`
	`text` TEXT NOT NULL, -- `full text of article`
    `user_id` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

CREATE TABLE users_articles (
    `id` SERIAL NOT NULL UNIQUE AUTO_INCREMENT,
    `user_id` INT NOT NULL, -- `user`
    `article_id` INT NOT NULL, -- `that user's articles`
    PRIMARY KEY (`id`),
    INDEX (`user_id`),
    INDEX (`article_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE
);
