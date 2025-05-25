/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 80037
 Source Host           : 192.168.61.2:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80037
 File Encoding         : 65001

 Date: 22/05/2025 10:05:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for nyx_user_role
-- ----------------------------
DROP TABLE IF EXISTS `nyx_user_role`;
CREATE TABLE `nyx_user_role`  (
  `id` int NOT NULL,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `roles` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of nyx_user_role
-- ----------------------------

-- ----------------------------
-- Table structure for nyx_users_info
-- ----------------------------
DROP TABLE IF EXISTS `nyx_users_info`;
CREATE TABLE `nyx_users_info`  (
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户uuid',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户名字',
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户电话号码',
  `introduction` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '用户介绍',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '提币地址',
  PRIMARY KEY (`uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of nyx_users_info
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
