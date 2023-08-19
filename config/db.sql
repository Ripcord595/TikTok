/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80100
 Source Host           : localhost:3306
 Source Schema         : TikTok

 Target Server Type    : MySQL
 Target Server Version : 80100
 File Encoding         : 65001

 Date: 03/08/2023 01:09:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  `comment_text` varchar(255) NOT NULL,
  `cancel_comment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0表示已评论，1表示取消评论',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `videoIndex` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1098 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论列表';

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `follower_id` bigint NOT NULL,
  `cancel_follow` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0表示关注，1则表示未关注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userIdtoFollowerIdIndex` (`user_id`,`follower_id`) USING BTREE,
  KEY `followerIdIndex` (`follower_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1302 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='关注列表';

-- ----------------------------
-- Records of follow
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for like
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  `cancle_like` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userIdtoVideoIdIndex` (`user_id`,`video_id`) USING BTREE,
  KEY `userIdIndex` (`user_id`) USING BTREE,
  KEY `videoIdIndex` (`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='点赞列表';

-- ----------------------------
-- Records of like
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  username varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name_password_index` (username,`password`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表\n';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `author_id` bigint NOT NULL,
  `play_url` varchar(255) NOT NULL,
  `cover_url` varchar(255) NOT NULL,
  `publish_time` datetime NOT NULL,
  `title` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `time` (`publish_time`) USING BTREE,
  KEY `author` (`author_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8mb3 COMMENT='视频表';

-- ----------------------------
-- Records of video
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Procedure structure for addFollowRelation
-- ----------------------------
DROP PROCEDURE IF EXISTS `addFollowRelation`;
delimiter ;;
CREATE PROCEDURE `addFollowRelation`(IN user_id INT, IN follower_id INT)
BEGIN
    -- 声明记录个数变量。
    DECLARE cnt INT DEFAULT 0;
    -- 获取记录个数变量。
    SELECT COUNT(1) FROM `follow` WHERE `user_id` = user_id AND `follower_id` = follower_id INTO cnt;
    -- 判断是否已经存在该记录，并做出相应的插入关系、更新关系动作。
    -- 插入操作。
    IF cnt = 0 THEN
        INSERT INTO `follow` (`user_id`, `follower_id`, `cancel_follow`) VALUES (user_id, follower_id, 0);
    END IF;
    -- 更新操作。
    IF cnt != 0 THEN
        UPDATE `follow` SET `cancel_follow` = 0 WHERE `user_id` = user_id AND `follower_id` = follower_id;
    END IF;
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for add_comment
-- ----------------------------
DROP PROCEDURE IF EXISTS `add_comment`;
delimiter ;;
CREATE PROCEDURE `add_comment`(IN user_id INT, IN video_id INT, IN comment_text VARCHAR(255))
BEGIN
    -- 插入评论记录
    INSERT INTO `comment` (`user_id`, `video_id`, `comment_text`, `create_time`)
    VALUES (user_id, video_id, comment_text, NOW());
    
    -- 返回评论ID作为结果
    SELECT LAST_INSERT_ID() AS comment_id;
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for data_statistics
-- ----------------------------
DROP PROCEDURE IF EXISTS `data_statistics`;
delimiter ;;
CREATE PROCEDURE `data_statistics`()
BEGIN
    -- 查询用户关注数和视频播放量
    SELECT
        (SELECT COUNT(*) FROM `follow`) AS total_followers,
        (SELECT SUM(play_count) FROM `video`) AS total_play_count;
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for delFollowRelation
-- ----------------------------
DROP PROCEDURE IF EXISTS `delFollowRelation`;
delimiter ;;
CREATE PROCEDURE `delFollowRelation`(IN user_id INT, IN follower_id INT)
BEGIN
    -- 定义记录个数变量，记录是否存在此关系，默认没有关系。
    DECLARE cnt INT DEFAULT 0;
    -- 查看是否之前有关系。
    SELECT COUNT(1) FROM `follow` WHERE `user_id` = user_id AND `follower_id` = follower_id INTO cnt;
    -- 有关系，则需要 update cancel_follow = 1，使其关系无效。
    IF cnt = 1 THEN
        UPDATE `follow` SET `cancel_follow` = 1 WHERE `user_id` = user_id AND `follower_id` = follower_id;
    END IF;
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for like_video
-- ----------------------------
DROP PROCEDURE IF EXISTS `like_video`;
delimiter ;;
CREATE PROCEDURE `like_video`(IN user_id INT, IN video_id INT)
BEGIN
    -- 声明变量用于验证是否已经点赞
    DECLARE like_count INT;
    
    -- 查询是否已经点赞
    SELECT COUNT(*) INTO like_count FROM `like` WHERE `user_id` = user_id AND `video_id` = video_id;
    
    -- 如果已经点赞，则取消点赞
    IF like_count > 0 THEN
        DELETE FROM `like` WHERE `user_id` = user_id AND `video_id` = video_id;
    ELSE
        -- 否则，进行点赞
        INSERT INTO `like` (`user_id`, `video_id`, `cancel_like`) VALUES (user_id, video_id, 0);
    END IF;
    
    -- 返回点赞状态（0 表示已点赞，1 表示取消点赞）
    SELECT IF(like_count > 0, 1, 0) AS like_status;
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for log_event
-- ----------------------------
DROP PROCEDURE IF EXISTS `log_event`;
delimiter ;;
CREATE PROCEDURE `log_event`(IN event_type VARCHAR(255), IN event_description TEXT)
BEGIN
    -- 插入日志记录
    INSERT INTO `log` (`event_type`, `event_description`, `event_time`)
    VALUES (event_type, event_description, NOW());
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for user_registration
-- ----------------------------
DROP PROCEDURE IF EXISTS `user_registration`;
delimiter ;;
CREATE PROCEDURE `user_registration`(IN username VARCHAR(255), IN userpassword VARCHAR(255))
BEGIN
    -- 声明变量用于验证用户名是否唯一
    DECLARE user_count INT;
    
    -- 查询用户名是否已经存在
    SELECT COUNT(*) INTO user_count FROM `user` WHERE username = username;
    
    -- 如果用户名已存在，则抛出错误
    IF user_count > 0 THEN
        SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = '用户名已存在';
    END IF;
    
    -- 否则，进行用户注册逻辑，将密码加密后存储
    INSERT INTO `user` (username, `password`) VALUES (username, SHA2(userpassword, 256));
    
    -- 返回注册成功信息
    SELECT '注册成功' AS message;
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for video_upload
-- ----------------------------
DROP PROCEDURE IF EXISTS `video_upload`;
delimiter ;;
CREATE PROCEDURE `video_upload`(IN author_id INT, IN play_url VARCHAR(255), IN cover_url VARCHAR(255), IN title VARCHAR(255))
BEGIN
    -- 生成视频ID，假设视频ID使用自增主键
    DECLARE video_id INT;
    
    -- 插入视频记录
    INSERT INTO `video` (`author_id`, `play_url`, `cover_url`, `publish_time`, `title`)
    VALUES (author_id, play_url, cover_url, NOW(), title);
    
    -- 获取生成的视频ID
    SET video_id = LAST_INSERT_ID();
    
    -- 返回视频ID作为结果
    SELECT video_id AS video_id;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
