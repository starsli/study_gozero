-- Drop database if exists
DROP DATABASE IF EXISTS `user_db`;

-- Create database
CREATE DATABASE IF NOT EXISTS `user_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- Use database
USE `user_db`;

DROP TABLE IF EXISTS `t_relation`;
CREATE TABLE `t_relation` (
  `user_id` VARCHAR(64) NOT NULL COMMENT '用户ID',
  `uid` BIGINT NOT NULL COMMENT '主键',
  `state` TINYINT NOT NULL COMMENT '注册状态',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  INDEX `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Drop table if exists
DROP TABLE IF EXISTS `t_user_info`;
-- Create user table
CREATE TABLE `t_user_info` (
  `uid` BIGINT NOT NULL COMMENT '主键',
  `user_id` VARCHAR(64) NOT NULL COMMENT '用户ID',
  `name` VARCHAR(64) NOT NULL COMMENT '用户名',
  `age` INTEGER NOT NULL COMMENT '年龄',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  INDEX `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `t_account`;

CREATE TABLE `t_account` (
  `uid` BIGINT NOT NULL COMMENT '主键',
  `user_id` VARCHAR(64) NOT NULL COMMENT '用户ID',
  `balance` BIGINT NOT NULL COMMENT '余额',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  INDEX `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `t_uid_segment`;

CREATE TABLE `t_uid_segment` (
  `id` BIGINT NOT NULL COMMENT '主键',
  `uid_max` BIGINT NOT NULL COMMENT '已经使用的UID最大值',
  `step` BIGINT NOT NULL COMMENT '步长',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Init test data
INSERT INTO `t_uid_segment` (`id`, `uid_max`, `step`) VALUES (1, 10000000, 1);
-- linux:  mysql -h 127.0.0.1 -P 3307 -u root -p123456 < user_init.sql
-- windows: Get-Content -Encoding UTF8 user_init.sql | mysql -h 127.0.0.1 -P 3307 -u root -p123456