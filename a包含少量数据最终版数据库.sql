/*
 Navicat MySQL Data Transfer

 Source Server         : New Connection
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : mytiktok

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 13/06/2022 14:56:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` int NOT NULL,
  `video_id` int NULL DEFAULT NULL,
  `user_id` int NULL DEFAULT NULL,
  `content` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_date` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, 1, 1, '1', '2022-06-03 21:42:19');
INSERT INTO `comments` VALUES (2, 1, 1, '222', '2022-06-01 21:42:13');
INSERT INTO `comments` VALUES (3, 1, 1, '第4条评论', '2022-06-13 11:08:19');
INSERT INTO `comments` VALUES (4, 1, 1, '第4条评论', '2022-06-13 11:08:48');
INSERT INTO `comments` VALUES (6, 1, 1, '第4条评论', '2022-06-13 11:39:52');

-- ----------------------------
-- Table structure for favoritesqls
-- ----------------------------
DROP TABLE IF EXISTS `favoritesqls`;
CREATE TABLE `favoritesqls`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `video_id` int NULL DEFAULT NULL,
  `user_id` int NULL DEFAULT NULL,
  `user_token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favoritesqls
-- ----------------------------
INSERT INTO `favoritesqls` VALUES (1, 1, 10, NULL, NULL);
INSERT INTO `favoritesqls` VALUES (3, 1, 10, NULL, NULL);
INSERT INTO `favoritesqls` VALUES (7, 1, 10, NULL, NULL);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `fans_counts` int NULL DEFAULT 0 COMMENT '我的粉丝数量',
  `follow_counts` int NULL DEFAULT 0 COMMENT '我关注的人总数',
  `receive_like_counts` int NULL DEFAULT 0 COMMENT '我接受到的赞美/收藏 的数量',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '123', '123456789', 111, 11, 1);
INSERT INTO `users` VALUES (3, '001', '123456789', 0, 0, 0);
INSERT INTO `users` VALUES (4, '002', '123456789', 0, 0, 0);
INSERT INTO `users` VALUES (5, '004', '123456789', 0, 0, 0);
INSERT INTO `users` VALUES (6, '005', '123456789', 0, 0, 0);
INSERT INTO `users` VALUES (7, '006', '123456789', 0, 0, 0);
INSERT INTO `users` VALUES (8, '1', '123445567', 0, 0, 0);
INSERT INTO `users` VALUES (9, '12', '122334444', 0, 0, 0);
INSERT INTO `users` VALUES (10, '007', '123456789', 0, 0, 0);
INSERT INTO `users` VALUES (11, '008', '123456789', 0, 0, 0);

-- ----------------------------
-- Table structure for users_like_videos
-- ----------------------------
DROP TABLE IF EXISTS `users_like_videos`;
CREATE TABLE `users_like_videos`  (
  `id` int NOT NULL,
  `user_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户',
  `video_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_video_rel`(`user_id`, `video_id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户喜欢的/赞过的视频' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users_like_videos
-- ----------------------------
INSERT INTO `users_like_videos` VALUES (1, '1', '1');
INSERT INTO `users_like_videos` VALUES (2, '3', '1');

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '发布者id',
  `video_title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频标题',
  `video_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频存放的路径',
  `cover_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频封面图',
  `like_counts` bigint NULL DEFAULT 0 COMMENT '喜欢/赞美的数量',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `favorite_count` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '视频信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, '1', '11', '', NULL, 11, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
