/*
 Navicat Premium Data Transfer

 Source Server         : article
 Source Server Type    : MySQL
 Source Server Version : 50716
 Source Host           : localhost:3306
 Source Schema         : goblin

 Target Server Type    : MySQL
 Target Server Version : 50716
 File Encoding         : 65001

 Date: 05/11/2019 19:57:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cards
-- ----------------------------
DROP TABLE IF EXISTS `cards`;
CREATE TABLE `cards`  (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `images` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `message` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `like` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cards
-- ----------------------------
INSERT INTO `cards` VALUES ('1572885440', 'test', 'dsfasdfass', '撒打发士大夫撒旦发射点发生', NULL, '1572523394');
INSERT INTO `cards` VALUES ('1572887401', 'testsdfsdsd', 'dsfasdfass', '撒打发士大夫撒旦发射点发生', NULL, '1572523394');
INSERT INTO `cards` VALUES ('1572890456', 'testsdfsdsd', 'dsfasdfass', '撒打发士大夫撒旦发射点发生', NULL, '1572890292');
INSERT INTO `cards` VALUES ('1572890458', 'testsdfsdsd', 'dsfasdfass', '撒打发士大夫撒旦发射点发生', NULL, '1572890292');

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo`  (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `photoer` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `userimage` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `message` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of userinfo
-- ----------------------------
INSERT INTO `userinfo` VALUES ('1572523394', 'test', '123456', NULL, '/avator/1b3094dc1746bd302845e1576193a8a03232f7db9f89a-CBB7i0_fw658.jpg', NULL);
INSERT INTO `userinfo` VALUES ('1572523418', '小负子', 'xfz123456', NULL, '/avator/1b3094dc1746bd302845e1576193a8a03232f7db9f89a-CBB7i0_fw658.jpg', NULL);
INSERT INTO `userinfo` VALUES ('1572582168', 'test', '123456', NULL, '', NULL);
INSERT INTO `userinfo` VALUES ('1572582381', 'asda', 'asdasdas', NULL, '', NULL);
INSERT INTO `userinfo` VALUES ('1572876520', 'zhang', '123', NULL, '', NULL);
INSERT INTO `userinfo` VALUES ('1572876521', 'zhang', '123', NULL, '', NULL);
INSERT INTO `userinfo` VALUES ('1572876627', 'zhang', '123', NULL, '', NULL);
INSERT INTO `userinfo` VALUES ('1572890292', 'zhg', '123', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', NULL);

SET FOREIGN_KEY_CHECKS = 1;
