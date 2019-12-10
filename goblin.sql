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

 Date: 10/12/2019 08:58:56
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
  `media` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cards
-- ----------------------------
INSERT INTO `cards` VALUES ('1572885440', 'test', '', '撒打发士大夫撒旦发射点发生是打发撒的发生大发生大傻傻的空间发货快乐四大上岛咖啡金卡束带结发开始电视剧阿飞开始大家福利卡圣撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是诞节福利卡撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是撒打发士大夫撒旦发射点发生是', NULL, '1572523394', '13850477-1-16.mp4');
INSERT INTO `cards` VALUES ('1572887401', 'testsdfsdsd', '8fdf8aa392eda2a57b543ccf7790a5d85e830245292f3-gKZ1Vq_fw658.jpg|9f3902df0f84b8b4f2888285df4a03692a8dd6e210da8-qfNAAg_fw658.jpg|57d1d3df3bcf94b6bba8b6a6d48fb980bd1e36c5166bc-nmkufP_fw658 (1).jpg', '撒打发士大夫撒旦发射点发生', NULL, '1572523394', '');
INSERT INTO `cards` VALUES ('1572890456', 'testsdfsdsd', '', '撒打发士大夫撒旦发射点发生', NULL, '1572890292', '13850477-1-16.mp4');
INSERT INTO `cards` VALUES ('1572890458', 'testsdfsdsd', '', '撒打发士大夫撒旦发射点发生', NULL, '1572890292', '13850477-1-16.mp4');
INSERT INTO `cards` VALUES ('1573107933', 'tzzzzzzzzzzzzz', '', 'zzzzzzzzzzzzzzzzzzzzz', NULL, '1572523394', '13850477-1-16.mp4');
INSERT INTO `cards` VALUES ('1573107943', 'tzzzzzzzzzzzzz', '', 'zzzzzzzzzzzzzzzzzzzzz', NULL, '1572523394', '13850477-1-16.mp4');
INSERT INTO `cards` VALUES ('1573795945', 'HelloWorld！！！', '8fdf8aa392eda2a57b543ccf7790a5d85e830245292f3-gKZ1Vq_fw658.jpg|9f3902df0f84b8b4f2888285df4a03692a8dd6e210da8-qfNAAg_fw658.jpg|57d1d3df3bcf94b6bba8b6a6d48fb980bd1e36c5166bc-nmkufP_fw658 (1).jpg', 'sdfsdfasdfasdfasd', NULL, '1572523394', '13850477-1-16.mp4');
INSERT INTO `cards` VALUES ('1573796646', '胜多负少', '', '士大夫撒地方', NULL, '1572523394', '13850477-1-16.mp4');
INSERT INTO `cards` VALUES ('1573809182', '是打发点', '', '撒的发生大发送到', NULL, '1572523394', '5505888-1-16.mp4');
INSERT INTO `cards` VALUES ('1573809239', '是打发傻傻的', '', '阿萨德法师法师', NULL, '1572523394', 'big_buck_bunny.mp4');
INSERT INTO `cards` VALUES ('1573809304', '是打阿达', '', '阿啊实打实大', NULL, '1572523394', 'big_buck_bunny.mp4');
INSERT INTO `cards` VALUES ('1575572927', '7啊是大', '0a3ab3b5f2a2c6e9e4f7a01ee964071ba23c512b6d1df-MAqB6T_fw658.jpg', '吱吱吱促销', NULL, '1572523394', '');
INSERT INTO `cards` VALUES ('1575572954', '试点范围为', '4fc31fba889b8bc2c3426cfe3c747943e0ce9720d0683-PG5OBS_fw658.jpg|4fc31fba889b8bc2c3426cfe3c747943e0ce9720d0683-PG5OBS_fw658.jpg|60324be8e79e6bf27bd4b580f4080e9a3316641ef50fd-qpfNyy_fw658.jpg', '阿斯蒂芬娃儿服务', NULL, '1572523394', '');
INSERT INTO `cards` VALUES ('1575573200', 'HelloWorld！！！呈现出v小东西', '72cb826d870439e864b4890fb53e2a8e8102278c197d7-EVUC6B_fw658.jpg|112cc6d4260e5c486009457b07eb7df973a896e7d447-1vRT1C_fw658.jpg|799c2968b24d00bbeadd608ae6970b488554388e31239-G5NVMA_fw658.jpg', '神无心，云无月', NULL, '1572523394', '');
INSERT INTO `cards` VALUES ('1575573268', '吱吱吱吱', '785593476d55619c4d8ff70979c3f31f93e510a1142f3-kuv5VW_fw658.jpg', '士大夫胜多负少我', NULL, '1572523394', '');

-- ----------------------------
-- Table structure for demand
-- ----------------------------
DROP TABLE IF EXISTS `demand`;
CREATE TABLE `demand`  (
  `id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `place` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `kind` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `message` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demand
-- ----------------------------
INSERT INTO `demand` VALUES ('1575564316', '川北医学院', '旅行', 'sdfsdaf sdfsadfsadfsadfsdfas sdfasdfas', '1572523394');
INSERT INTO `demand` VALUES ('1575564332', '川北医学院', '旅行', 'sdfsdaf sdfsadfsadfsadfsdfas sdfasdfas', '1572523394');
INSERT INTO `demand` VALUES ('1575564364', '西华师范大学', '街拍', 'lkasdgklasdjflksdajflkasdjlfkasf', '1572523394');
INSERT INTO `demand` VALUES ('1575564381', '西华师范大学', '校园风', '健康快乐看来很快加快了可怜见立刻', '1572523394');
INSERT INTO `demand` VALUES ('1575564386', '西华师范大学', '校园风', '健康快乐看来很快加快了可怜见立刻', '1572523394');
INSERT INTO `demand` VALUES ('1575564388', '西华师范大学', '校园风', '健康快乐看来很快加快了可怜见立刻', '1572523394');

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
INSERT INTO `userinfo` VALUES ('1572523394', 'test', '123456', NULL, '/avator/1b3094dc1746bd302845e1576193a8a03232f7db9f89a-CBB7i0_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1572523418', '小负子', 'xfz123456', NULL, '/avator/1b3094dc1746bd302845e1576193a8a03232f7db9f89a-CBB7i0_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1572582168', 'test', '123456', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1572582381', 'asda', 'asdasdas', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1572876520', 'zhang', '123', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1572876521', 'zhang', '123', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1572876627', 'zhang', '123', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1572890292', 'zhg', '123', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1573200591', 'asasa', 'aaaaaaa', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1573200609', 'asasa', 'aaaaaaa', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1573200666', 'sdfas ', 'sadfasd', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1573200717', 'asdfas', 'asdfasfd', NULL, '/avator/0b7061bcf2f05ea41f8649938362286df9cd3dc7d437-GvxzBF_fw658.jpg', 'asdfasgasfsdfasdfasdfasfasdfa');
INSERT INTO `userinfo` VALUES ('1573200722', 'asdfas', 'asdfasfd', NULL, NULL, 'asdfasgasfsdfasdfasdfasfasdfa');

SET FOREIGN_KEY_CHECKS = 1;
