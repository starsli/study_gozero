CREATE TABLE `t_user` (
  `id` BIGINT AUTO_INCREMENT COMMENT '主键',
  `name` VARCHAR(64) NOT NULL COMMENT '用户名',
  `age` TINYINT NOT NULL COMMENT '年龄',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_update_time` (`update_time`),
  INDEX `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- goctl model mysql ddl -src ./model/mysql/user.sql -dir ./model/mysql -c
-- -c：开启缓存（redis，可选，不加则无缓存）
