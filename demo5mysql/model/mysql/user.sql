
CREATE TABLE `t_relation` (
  `user_id` VARCHAR(64) NOT NULL COMMENT '用户ID',
  `uid` BIGINT NOT NULL COMMENT '主键',
  `state` TINYINT NOT NULL COMMENT '关联状态',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  INDEX `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `t_user_info` (
  `uid` BIGINT NOT NULL COMMENT '主键',
  `user_id` VARCHAR(64) NOT NULL COMMENT '用户ID',
  `password` VARCHAR(128) NOT NULL COMMENT '支付密码',
  `name` VARCHAR(128) NOT NULL COMMENT '姓名',
  `gender` TINYINT NOT NULL COMMENT '性别',
  `age` SMALLINT NOT NULL COMMENT '年龄',
  `address` VARCHAR(128) NOT NULL COMMENT '地址',
  `phone` VARCHAR(128) NOT NULL COMMENT '手机号',
  `email` VARCHAR(128) NOT NULL COMMENT '邮箱',
  `id_type` TINYINT NOT NULL COMMENT '身份证类型',
  `id_card` VARCHAR(128) NOT NULL COMMENT '身份证号',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  INDEX `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `t_account` (
  `uid` BIGINT NOT NULL COMMENT '主键',
  `user_id` VARCHAR(64) NOT NULL COMMENT '用户ID',
  `balance` BIGINT NOT NULL COMMENT '余额',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  INDEX `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户流水表
CREATE TABLE `t_account_flow` (
  `id` BIGINT AUTO_INCREMENT COMMENT '主键',
  `uid` BIGINT NOT NULL COMMENT '用户UID',
  `counterparty_id` BIGINT NOT NULL COMMENT '对方用户UID',
  `user_id` VARCHAR(64) NOT NULL COMMENT '用户ID',
  `flow_id` VARCHAR(64) NOT NULL COMMENT '流水ID',
  `flow_type` TINYINT NOT NULL COMMENT '流水类型',
  `biz_type` TINYINT NOT NULL COMMENT '业务类型',
  `amount` BIGINT NOT NULL COMMENT '金额',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_uid_direct_type` (`uid`,`flow_type`,`flow_id`),
  INDEX `idx_flow_id` (`flow_id`),
  INDEX `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `t_uid_segment`;

CREATE TABLE `t_uid_segment` (
  `id` BIGINT NOT NULL COMMENT '主键',
  `uid_max` BIGINT NOT NULL COMMENT '已使用的最大用户ID',
  `step` BIGINT NOT NULL COMMENT '步长',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- goctl model mysql ddl -src user.sql -dir .
-- -c：开启缓存（redis，可选，不加则无缓存）
