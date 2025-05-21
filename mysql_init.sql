/*
 Navicat Premium Data Transfer

 Source Server         : go_test
 Source Server Type    : MySQL
 Source Server Version : 80037
 Source Host           : 192.168.2.7:3306
 Source Schema         : nyx

 Target Server Type    : MySQL
 Target Server Version : 80037
 File Encoding         : 65001

 Date: 21/05/2025 22:37:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of nyx_users_info
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
