CREATE TABLE articles (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	`title` VARCHAR(100) NOT NULL, -- `title of article`
	`announce` VARCHAR(255), -- `anouncement of article`
	`text` TEXT NOT NULL -- `full text of article`
);
