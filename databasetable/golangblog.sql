/*
 Navicat Premium Data Transfer

 Source Server         : 本机数据库
 Source Server Type    : MySQL
 Source Server Version : 50562
 Source Host           : localhost:3306
 Source Schema         : golangblog

 Target Server Type    : MySQL
 Target Server Version : 50562
 File Encoding         : 65001

 Date: 31/08/2020 17:01:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for db_articlesort
-- ----------------------------
DROP TABLE IF EXISTS `db_articlesort`;
CREATE TABLE `db_articlesort`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_sort` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `general_sort` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of db_articlesort
-- ----------------------------
INSERT INTO `db_articlesort` VALUES (1, 'MySQL', '1');
INSERT INTO `db_articlesort` VALUES (2, 'Redis', '1');
INSERT INTO `db_articlesort` VALUES (3, 'Nginx', '1');
INSERT INTO `db_articlesort` VALUES (4, 'Vue', '1');
INSERT INTO `db_articlesort` VALUES (5, 'GoLang', '1');
INSERT INTO `db_articlesort` VALUES (6, 'Java', '1');
INSERT INTO `db_articlesort` VALUES (7, '旅游', '2');
INSERT INTO `db_articlesort` VALUES (8, '日常', '2');
INSERT INTO `db_articlesort` VALUES (9, '生活', '2');
INSERT INTO `db_articlesort` VALUES (10, '工作', '3');

-- ----------------------------
-- Table structure for db_general_sort
-- ----------------------------
DROP TABLE IF EXISTS `db_general_sort`;
CREATE TABLE `db_general_sort`  (
  `id` int(32) NOT NULL,
  `general_sort` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of db_general_sort
-- ----------------------------
INSERT INTO `db_general_sort` VALUES (1, '学习');
INSERT INTO `db_general_sort` VALUES (2, '生活');
INSERT INTO `db_general_sort` VALUES (3, '经历');

-- ----------------------------
-- Table structure for db_sex
-- ----------------------------
DROP TABLE IF EXISTS `db_sex`;
CREATE TABLE `db_sex`  (
  `id` int(100) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of db_sex
-- ----------------------------
INSERT INTO `db_sex` VALUES (1, '女');
INSERT INTO `db_sex` VALUES (2, '男');

-- ----------------------------
-- Table structure for tb_article
-- ----------------------------
DROP TABLE IF EXISTS `tb_article`;
CREATE TABLE `tb_article`  (
  `id` int(50) NOT NULL AUTO_INCREMENT,
  `article_sort` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章分类，外键',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文章标题',
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文章内容',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tb_article
-- ----------------------------
INSERT INTO `tb_article` VALUES (12, '5', '测试', '测试内容', '2020-08-28 11:33:28', '2020-08-31 14:10:41');
INSERT INTO `tb_article` VALUES (13, '1', '测试', '测试内容', '2020-08-28 11:33:28', '2020-08-31 11:23:19');
INSERT INTO `tb_article` VALUES (15, '7', '测试', '测试内容', '2020-08-28 11:33:28', '2020-08-31 11:37:11');
INSERT INTO `tb_article` VALUES (16, '7', '测试', '测试内', '2020-08-31 14:35:43', '2020-08-31 14:35:59');

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `id` int(64) NOT NULL AUTO_INCREMENT,
  `openid` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_login_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` datetime NOT NULL,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES (1, NULL, '曾帅', '02', '18275507824', '18275507824', '12345678', '2020-07-17 09:24:38', '2020-08-17 08:43:27');
INSERT INTO `tb_user` VALUES (2, NULL, '小晓', '01', '18652732938', '18652732938', '12345678', '2020-08-13 14:05:31', '2020-08-17 08:43:56');

SET FOREIGN_KEY_CHECKS = 1;
