SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
                             `id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                             `father_comment_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '父评论id',
                             `to_user_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '被评论的用户id',
                             `video_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频id',
                             `from_user_id` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '留言者，评论的用户id',
                             `comment` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '评论内容',
                             `create_time` datetime NOT NULL,
                             PRIMARY KEY (`id`) USING BTREE,
                             INDEX `to_user_id`(`to_user_id`) USING BTREE,
                             INDEX `father_comment_id`(`father_comment_id`) USING BTREE,
                             INDEX `video_id`(`video_id`) USING BTREE,
                             INDEX `from_user_id`(`from_user_id`) USING BTREE,
                             CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`to_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                             CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`father_comment_id`) REFERENCES `comments` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                             CONSTRAINT `comments_ibfk_3` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                             CONSTRAINT `comments_ibfk_4` FOREIGN KEY (`from_user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '视频评论表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comments
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                          `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
                          `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
                          `face_image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '我的头像，如果没有默认给一张',
                          `nickname` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
                          `fans_counts` int NULL DEFAULT 0 COMMENT '我的粉丝数量',
                          `follow_counts` int NULL DEFAULT 0 COMMENT '我关注的人总数',
                          `receive_like_counts` int NULL DEFAULT 0 COMMENT '我接受到的赞美/收藏 的数量',
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `id`(`id`) USING BTREE,
                          UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
-- ----------------------------
-- Table structure for users_fans
-- ----------------------------
DROP TABLE IF EXISTS `users_fans`;
CREATE TABLE `users_fans`  (
                               `id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                               `user_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户',
                               `fan_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '粉丝',
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `user_id`(`user_id`, `fan_id`) USING BTREE,
                               INDEX `fan_id`(`fan_id`) USING BTREE,
                               CONSTRAINT `users_fans_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                               CONSTRAINT `users_fans_ibfk_2` FOREIGN KEY (`fan_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户粉丝关联关系表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users_fans
-- ----------------------------

-- ----------------------------
-- Table structure for users_like_videos
-- ----------------------------
DROP TABLE IF EXISTS `users_like_videos`;
CREATE TABLE `users_like_videos`  (
                                      `id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                                      `user_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户',
                                      `video_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频',
                                      PRIMARY KEY (`id`) USING BTREE,
                                      UNIQUE INDEX `user_video_rel`(`user_id`, `video_id`) USING BTREE,
                                      INDEX `video_id`(`video_id`) USING BTREE,
                                      CONSTRAINT `users_like_videos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                                      CONSTRAINT `users_like_videos_ibfk_2` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户喜欢的/赞过的视频' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users_like_videos
-- ----------------------------

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
                           `id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                           `user_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '发布者id',
                           `video_desc` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频描述',
                           `video_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '视频存放的路径',
                           `video_seconds` float(6, 2) NULL DEFAULT NULL COMMENT '视频秒数',
  `video_width` int NULL DEFAULT NULL COMMENT '视频宽度',
  `video_height` int NULL DEFAULT NULL COMMENT '视频高度',
  `cover_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频封面图',
  `like_counts` bigint NOT NULL DEFAULT 0 COMMENT '喜欢/赞美的数量',
  `status` int NOT NULL COMMENT '视频状态：\r\n1、发布成功\r\n2、禁止播放，管理员操作',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `videos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '视频信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of videos
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;