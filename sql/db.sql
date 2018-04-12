CREATE DATABASE friendly_reminder;

CREATE TABLE friendly_reminder.`notice` (
  `notice_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '提醒id',
  `from_user_id` bigint(20) NOT NULL COMMENT '创建人用户id',
  `to_user_id` bigint(20) NOT NULL COMMENT '接受人用户id',
  `notice_title` varchar(32) NOT NULL COMMENT '提醒标题',
  `notice_content` varchar(255) NOT NULL COMMENT '提醒详细内容',
  `notice_time` bigint(64) NOT NULL COMMENT '提醒时间',
  PRIMARY KEY (`notice_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通知表';

CREATE TABLE friendly_reminder.`user` (
  `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_name` varchar(8) NOT NULL COMMENT '用户名',
  `user_phone` varchar(11) NOT NULL COMMENT '用户手机号码',
  `user_password` char(32) NOT NULL COMMENT '用户密码',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_name` (`user_name`),
  KEY `user_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='用户表';

CREATE TABLE friendly_reminder.`user_user` (
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `friend_id` bigint(20) NOT NULL COMMENT '好友id',
  `is_agree` int(1) NOT NULL DEFAULT '0',
  UNIQUE KEY `user_id` (`user_id`,`friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户关系表';
