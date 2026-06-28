-- Drop database if exists
DROP DATABASE IF EXISTS `user_db`;

-- Create database
CREATE DATABASE IF NOT EXISTS `user_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- Use database
USE `user_db`;

-- Drop table if exists
DROP TABLE IF EXISTS `t_user`;

-- Create user table
CREATE TABLE IF NOT EXISTS `t_user` (
  `id` BIGINT AUTO_INCREMENT COMMENT '主键',
  `name` VARCHAR(64) NOT NULL COMMENT '用户名',
  `age` TINYINT NOT NULL COMMENT '年龄',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_update_time` (`update_time`),
  INDEX `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- Init test data
INSERT INTO `t_user` (`name`, `age`) VALUES ('zhangsan', 25);
INSERT INTO `t_user` (`name`, `age`) VALUES ('lisi', 30);
INSERT INTO `t_user` (`name`, `age`) VALUES ('wangwu', 28);

-- linux:  mysql -h 127.0.0.1 -P 3307 -u root -p123456 < user_init.sql
-- windows: Get-Content -Encoding UTF8 user_init.sql | mysql -h 127.0.0.1 -P 3307 -u root -p123456