-- name: create-database
CREATE DATABASE IF NOT EXISTS `meterdb`;

-- name: switch-to-database
USE `meterdb`;

-- name: create-meter-table
CREATE TABLE IF NOT EXISTS `meter` (
  `uid` INT(10) NOT NULL AUTO_INCREMENT,
  `meter_id` VARCHAR(64) NOT NULL,
  `meter_type` VARCHAR(64) NOT NULL,
  `master_key` VARCHAR(64) NULL,
  `data_key` VARCHAR(64) NULL,
  `create_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`)
);

-- name: create-meter
INSERT INTO meter (meter_id, meter_type) VALUES(?, ?)

-- name: find-one-meter-by-meter_id
SELECT meter_id FROM meter WHERE meter_id = ? LIMIT 1

-- name: find-one-meter-mk-by-meter_id
SELECT master_key FROM meter WHERE meter_id = ? LIMIT 1

-- name: find-one-meter-dk-by-meter_id
SELECT data_key FROM meter WHERE meter_id = ? LIMIT 1

-- name: update-meter-mk-by-meter_id
UPDATE meter SET master_key = ? WHERE meter_id = ?







-- name: create-userdetail
CREATE TABLE `userdetail` (
  `uid` INT(10) NOT NULL DEFAULT '0',
  `intro` TEXT NULL,
  `profile` TEXT NULL,
  PRIMARY KEY (`uid`)
)