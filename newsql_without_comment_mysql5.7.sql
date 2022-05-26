/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 26/05/2022 13:47:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(64) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `password` varchar(64) NOT NULL COMMENT '密码',
  `fans_counts` int(11) DEFAULT '0' COMMENT '我的粉丝数量',
  `follow_counts` int(11) DEFAULT '0' COMMENT '我关注的人总数',
  `receive_like_counts` int(11) DEFAULT '0' COMMENT '我接受到的赞美/收藏 的数量',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `id` (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for users_fans
-- ----------------------------
DROP TABLE IF EXISTS `users_fans`;
CREATE TABLE `users_fans` (
  `id` int(11) NOT NULL,
  `user_id` varchar(64) NOT NULL COMMENT '用户',
  `fan_id` varchar(64) NOT NULL COMMENT '粉丝',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id` (`user_id`,`fan_id`) USING BTREE,
  KEY `fan_id` (`fan_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户粉丝关联关系表';

-- ----------------------------
-- Table structure for users_like_videos
-- ----------------------------
DROP TABLE IF EXISTS `users_like_videos`;
CREATE TABLE `users_like_videos` (
  `id` int(11) NOT NULL,
  `user_id` varchar(64) NOT NULL COMMENT '用户',
  `video_id` varchar(64) NOT NULL COMMENT '视频',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_video_rel` (`user_id`,`video_id`) USING BTREE,
  KEY `video_id` (`video_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户喜欢的/赞过的视频';

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` int(11) NOT NULL,
  `user_id` varchar(64) NOT NULL COMMENT '发布者id',
  `video_desc` varchar(128) DEFAULT NULL COMMENT '视频描述',
  `video_path` varchar(255) NOT NULL COMMENT '视频存放的路径',
  `cover_path` varchar(255) DEFAULT NULL COMMENT '视频封面图',
  `like_counts` bigint(20) NOT NULL DEFAULT '0' COMMENT '喜欢/赞美的数量',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='视频信息表';

SET FOREIGN_KEY_CHECKS = 1;
