CREATE TABLE `package_memorizing` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint unsigned NOT NULL COMMENT ' 用户 ID',
  `package_type` bigint unsigned NOT NULL COMMENT '词包类型',
  `package_id` bigint unsigned NOT NULL COMMENT '词包ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid_package_id` (`uid`, `package_id`),
  KEY `package_id` (`package_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100001 DEFAULT CHARSET=utf8mb4;
