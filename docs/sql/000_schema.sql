-- Create Table Users
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL DEFAULT '',
  `username` VARCHAR(255) NOT NULL DEFAULT '',
  `password` VARCHAR(255) NOT NULL DEFAULT '',
  `display_name` VARCHAR(255) NOT NULL DEFAULT '',

  -- Utility columns
  `status` SMALLINT NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(255),
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(255),
  `deleted_at`TIMESTAMP,
  `deleted_by` VARCHAR(255),
  PRIMARY KEY (`id`),
  UNIQUE (`username`)
) ENGINE = INNODB COMMENT='User table';

-- Create Table Urls
-- For example from url shortener
DROP TABLE IF EXISTS `url`;
CREATE TABLE IF NOT EXISTS `url` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `fk_user_id` INT COMMENT 'Foreign Key to User ID',
  `original_url` VARCHAR(255) NOT NULL DEFAULT '',
  `shorten_url` VARCHAR(255) NOT NULL DEFAULT '',
  `visit` INT NOT NULL DEFAULT '0',

  -- Utility columns
  `status` SMALLINT NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(255),
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(255),
  `deleted_at`TIMESTAMP,
  `deleted_by` VARCHAR(255),
  PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT='url table';
