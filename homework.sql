/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80011
 Source Host           : localhost:3306
 Source Schema         : homework

 Target Server Type    : MySQL
 Target Server Version : 80011
 File Encoding         : 65001

 Date: 09/01/2021 22:48:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `telephone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `authority` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `telephone`(`telephone`) USING BTREE,
  INDEX `idx_admins_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES (1, '2021-01-03 10:38:30', '2021-01-03 10:38:30', NULL, 'XY', '18936660888', '$2a$10$IAzhHYvE9fQR/AGxmO9T.OxlBxqYMNJZYOUbGxnIi52Xoz16Lp12K', '1');
INSERT INTO `admins` VALUES (2, '2021-01-03 10:55:38', '2021-01-03 10:55:38', NULL, 'ARIA', '18936660688', '$2a$10$k9Jj3KSJqOiWkIBLAaFI2.upEPyanpupTteEHurtX/veaVZ6lSvV2', '0');
INSERT INTO `admins` VALUES (3, '2021-01-03 23:39:44', '2021-01-03 23:39:44', NULL, 'A', '18936660788', '$2a$10$LWULyOXHCs18IgaPLVgxneuJQx07hOuURudSTvHEz0clh1YONlLBu', '1');
INSERT INTO `admins` VALUES (4, '2021-01-07 15:12:44', '2021-01-07 15:12:44', NULL, 'XUYAO', '18932548795', '$2a$10$MN.XXpJO4ro1RfZxD1Ic3usB0n8xnhQN5QGSGFEU/J4GM.d.Mb3ZS', '0');

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `ProductID` int(30) NULL DEFAULT NULL,
  `IP` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Operation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `Time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 62 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of log
-- ----------------------------
INSERT INTO `log` VALUES (6, '1', 1, '169.254.157.217', 'delete', '2021/01/03 22:45:38');
INSERT INTO `log` VALUES (7, '1', 17, '169.254.157.217', 'delete', '2021/01/03 23:07:33');
INSERT INTO `log` VALUES (8, '0', 18, '169.254.157.217', 'delete', '2021/01/03 23:11:46');
INSERT INTO `log` VALUES (9, '0', 16, '169.254.157.217', 'delete', '2021/01/03 23:40:01');
INSERT INTO `log` VALUES (13, '胥耀', 1, '169.254.157.217', 'delete', '2021/01/03 23:59:10');
INSERT INTO `log` VALUES (14, '胥耀', 0, '169.254.157.217', 'update', '2021/01/04 08:34:56');
INSERT INTO `log` VALUES (15, '胥耀', 3, '169.254.157.217', 'update', '2021/01/04 08:41:48');
INSERT INTO `log` VALUES (16, '胥耀', 0, '169.254.157.217', 'check', '2021/01/04 08:49:27');
INSERT INTO `log` VALUES (17, '胥耀', 0, '169.254.157.217', 'check', '2021/01/04 09:00:05');
INSERT INTO `log` VALUES (18, '胥耀', 0, '169.254.157.217', 'login', '2021/01/04 09:14:17');
INSERT INTO `log` VALUES (19, '胥耀', 0, '169.254.157.217', 'add product', '2021/01/04 09:25:24');
INSERT INTO `log` VALUES (20, '胥耀', 0, '169.254.157.217', 'login', '2021/01/04 09:25:24');
INSERT INTO `log` VALUES (22, '胥耀', 0, '169.254.157.217', 'login', '2021/01/04 10:42:12');
INSERT INTO `log` VALUES (23, '胥耀', 0, '169.254.157.217', 'login', '2021/01/06 15:29:08');
INSERT INTO `log` VALUES (24, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 15:29:15');
INSERT INTO `log` VALUES (25, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 15:29:20');
INSERT INTO `log` VALUES (26, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 15:47:51');
INSERT INTO `log` VALUES (27, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 15:47:59');
INSERT INTO `log` VALUES (28, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 15:52:30');
INSERT INTO `log` VALUES (29, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 15:52:37');
INSERT INTO `log` VALUES (30, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 16:00:50');
INSERT INTO `log` VALUES (31, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 16:00:55');
INSERT INTO `log` VALUES (32, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 16:33:00');
INSERT INTO `log` VALUES (33, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 16:33:08');
INSERT INTO `log` VALUES (34, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 16:33:41');
INSERT INTO `log` VALUES (35, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 16:33:46');
INSERT INTO `log` VALUES (36, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 16:42:13');
INSERT INTO `log` VALUES (37, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 16:53:39');
INSERT INTO `log` VALUES (38, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 16:53:43');
INSERT INTO `log` VALUES (39, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 20:44:44');
INSERT INTO `log` VALUES (40, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 20:44:52');
INSERT INTO `log` VALUES (41, '胥耀', 0, '169.254.157.217', 'login', '2021/01/06 21:12:52');
INSERT INTO `log` VALUES (42, '胥耀', 1, '169.254.157.217', 'check order', '2021/01/06 21:13:58');
INSERT INTO `log` VALUES (43, '胥耀', 1, '169.254.157.217', 'update order', '2021/01/06 21:14:04');
INSERT INTO `log` VALUES (44, '胥耀', 0, '169.254.157.217', 'login', '2021/01/06 21:14:04');
INSERT INTO `log` VALUES (45, '胥耀', 1, '169.254.157.217', 'check product', '2021/01/06 21:56:44');
INSERT INTO `log` VALUES (46, '胥耀', 1, '169.254.157.217', 'change product', '2021/01/06 21:56:50');
INSERT INTO `log` VALUES (47, 'XY', 1, '169.254.157.217', 'check product', '2021/01/07 15:20:27');
INSERT INTO `log` VALUES (48, 'XY', 1, '169.254.157.217', 'check product', '2021/01/07 15:20:59');
INSERT INTO `log` VALUES (49, 'XY', 0, '169.254.157.217', 'login', '2021/01/07 15:21:57');
INSERT INTO `log` VALUES (50, 'XY', 1, '169.254.157.217', 'check product', '2021/01/07 15:22:05');
INSERT INTO `log` VALUES (51, 'XY', 1, '169.254.157.217', 'check product', '2021/01/07 15:22:23');
INSERT INTO `log` VALUES (52, 'XY', 2, '169.254.157.217', 'delete product', '2021/01/07 15:23:14');
INSERT INTO `log` VALUES (53, 'XY', 3, '169.254.157.217', 'delete product', '2021/01/07 15:23:19');
INSERT INTO `log` VALUES (54, 'XY', 0, '169.254.157.217', 'login', '2021/01/07 15:30:23');
INSERT INTO `log` VALUES (55, 'XY', 0, '169.254.157.217', 'login', '2021/01/07 15:30:28');
INSERT INTO `log` VALUES (56, 'XY', 0, '169.254.157.217', 'login', '2021/01/07 15:30:29');
INSERT INTO `log` VALUES (57, 'XY', 51, '169.254.157.217', 'check order', '2021/01/07 15:30:52');
INSERT INTO `log` VALUES (58, 'XY', 51, '169.254.157.217', 'check order', '2021/01/07 15:31:31');
INSERT INTO `log` VALUES (59, 'ARIA', 0, '169.254.157.217', 'login', '2021/01/07 18:12:05');
INSERT INTO `log` VALUES (60, 'ARIA', 0, '169.254.157.217', 'login', '2021/01/07 18:12:05');
INSERT INTO `log` VALUES (61, 'XY', 0, '169.254.157.217', 'login', '2021/01/07 18:33:39');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `userID` int(255) NULL DEFAULT NULL,
  `productID` int(255) NULL DEFAULT NULL,
  `orderStatus` int(255) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 101 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (1, 0, 1, 1);
INSERT INTO `order` VALUES (2, 0, 1, 1);
INSERT INTO `order` VALUES (3, 0, 1, 1);
INSERT INTO `order` VALUES (4, 0, 1, 1);
INSERT INTO `order` VALUES (5, 0, 1, 1);
INSERT INTO `order` VALUES (6, 0, 1, 1);
INSERT INTO `order` VALUES (7, 0, 1, 1);
INSERT INTO `order` VALUES (8, 0, 1, 1);
INSERT INTO `order` VALUES (9, 0, 1, 1);
INSERT INTO `order` VALUES (10, 0, 1, 1);
INSERT INTO `order` VALUES (11, 0, 1, 1);
INSERT INTO `order` VALUES (12, 0, 1, 1);
INSERT INTO `order` VALUES (13, 0, 1, 1);
INSERT INTO `order` VALUES (14, 0, 1, 1);
INSERT INTO `order` VALUES (15, 0, 1, 1);
INSERT INTO `order` VALUES (16, 0, 1, 1);
INSERT INTO `order` VALUES (17, 0, 1, 1);
INSERT INTO `order` VALUES (18, 0, 1, 1);
INSERT INTO `order` VALUES (19, 0, 1, 1);
INSERT INTO `order` VALUES (20, 0, 1, 1);
INSERT INTO `order` VALUES (21, 0, 1, 1);
INSERT INTO `order` VALUES (22, 0, 1, 1);
INSERT INTO `order` VALUES (23, 0, 1, 1);
INSERT INTO `order` VALUES (24, 0, 1, 1);
INSERT INTO `order` VALUES (25, 0, 1, 1);
INSERT INTO `order` VALUES (26, 0, 1, 1);
INSERT INTO `order` VALUES (27, 0, 1, 1);
INSERT INTO `order` VALUES (28, 0, 1, 1);
INSERT INTO `order` VALUES (29, 0, 1, 1);
INSERT INTO `order` VALUES (30, 0, 1, 1);
INSERT INTO `order` VALUES (31, 0, 1, 1);
INSERT INTO `order` VALUES (32, 0, 1, 1);
INSERT INTO `order` VALUES (33, 0, 1, 1);
INSERT INTO `order` VALUES (34, 0, 1, 1);
INSERT INTO `order` VALUES (35, 0, 1, 1);
INSERT INTO `order` VALUES (36, 0, 1, 1);
INSERT INTO `order` VALUES (37, 0, 1, 1);
INSERT INTO `order` VALUES (38, 0, 1, 1);
INSERT INTO `order` VALUES (39, 0, 1, 1);
INSERT INTO `order` VALUES (40, 0, 1, 1);
INSERT INTO `order` VALUES (41, 0, 1, 1);
INSERT INTO `order` VALUES (42, 0, 1, 1);
INSERT INTO `order` VALUES (43, 0, 1, 1);
INSERT INTO `order` VALUES (44, 0, 1, 1);
INSERT INTO `order` VALUES (45, 0, 1, 1);
INSERT INTO `order` VALUES (46, 0, 1, 1);
INSERT INTO `order` VALUES (47, 0, 1, 1);
INSERT INTO `order` VALUES (48, 0, 1, 1);
INSERT INTO `order` VALUES (49, 0, 1, 1);
INSERT INTO `order` VALUES (50, 0, 1, 1);
INSERT INTO `order` VALUES (51, 0, 1, 1);
INSERT INTO `order` VALUES (52, 0, 1, 1);
INSERT INTO `order` VALUES (53, 0, 1, 1);
INSERT INTO `order` VALUES (54, 0, 1, 1);
INSERT INTO `order` VALUES (55, 0, 1, 1);
INSERT INTO `order` VALUES (56, 0, 1, 1);
INSERT INTO `order` VALUES (57, 0, 1, 1);
INSERT INTO `order` VALUES (58, 0, 1, 1);
INSERT INTO `order` VALUES (59, 0, 1, 1);
INSERT INTO `order` VALUES (60, 0, 1, 1);
INSERT INTO `order` VALUES (61, 0, 1, 1);
INSERT INTO `order` VALUES (62, 0, 1, 1);
INSERT INTO `order` VALUES (63, 0, 1, 1);
INSERT INTO `order` VALUES (64, 0, 1, 1);
INSERT INTO `order` VALUES (65, 0, 1, 1);
INSERT INTO `order` VALUES (66, 0, 1, 1);
INSERT INTO `order` VALUES (67, 0, 1, 1);
INSERT INTO `order` VALUES (68, 0, 1, 1);
INSERT INTO `order` VALUES (69, 0, 1, 1);
INSERT INTO `order` VALUES (70, 0, 1, 1);
INSERT INTO `order` VALUES (71, 0, 1, 1);
INSERT INTO `order` VALUES (72, 0, 1, 1);
INSERT INTO `order` VALUES (73, 0, 1, 1);
INSERT INTO `order` VALUES (74, 0, 1, 1);
INSERT INTO `order` VALUES (75, 0, 1, 1);
INSERT INTO `order` VALUES (76, 0, 1, 1);
INSERT INTO `order` VALUES (77, 0, 1, 1);
INSERT INTO `order` VALUES (78, 0, 1, 1);
INSERT INTO `order` VALUES (79, 0, 1, 1);
INSERT INTO `order` VALUES (80, 0, 1, 1);
INSERT INTO `order` VALUES (81, 0, 1, 1);
INSERT INTO `order` VALUES (82, 0, 1, 1);
INSERT INTO `order` VALUES (83, 0, 1, 1);
INSERT INTO `order` VALUES (84, 0, 1, 1);
INSERT INTO `order` VALUES (85, 0, 1, 1);
INSERT INTO `order` VALUES (86, 0, 1, 1);
INSERT INTO `order` VALUES (87, 0, 1, 1);
INSERT INTO `order` VALUES (88, 0, 1, 1);
INSERT INTO `order` VALUES (89, 0, 1, 1);
INSERT INTO `order` VALUES (90, 0, 1, 1);
INSERT INTO `order` VALUES (91, 0, 1, 1);
INSERT INTO `order` VALUES (92, 0, 1, 1);
INSERT INTO `order` VALUES (93, 0, 1, 1);
INSERT INTO `order` VALUES (94, 0, 1, 1);
INSERT INTO `order` VALUES (95, 0, 1, 1);
INSERT INTO `order` VALUES (96, 0, 1, 1);
INSERT INTO `order` VALUES (97, 0, 1, 1);
INSERT INTO `order` VALUES (98, 0, 1, 1);
INSERT INTO `order` VALUES (99, 0, 1, 1);
INSERT INTO `order` VALUES (100, 0, 1, 1);

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `productName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `productNum` int(255) NULL DEFAULT NULL,
  `productPrice` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `productOriginPrice` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `productImage` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `productUrl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (1, 'apple', 0, '250', '300', 'https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3262485291,1572432962&fm=26&gp=0.jpg', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `userName` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `passWord` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (4, '', 'A', '$2a$10$Y45X.Gu3cmb95wwpzS5v2uqrRV/hm2C7imytoBkENyF7wjqpewFIu');
INSERT INTO `user` VALUES (5, 'AB', 'AB', '$2a$10$uWDUAqrWckpL1/QJA2/ri.uf/Gxej3p52h./S7K3s1peNurSIzWxO');
INSERT INTO `user` VALUES (6, 'AAA', 'AAA', '$2a$10$tBjiyx0q0YZIb5SUJ/hGOuC6sZ/ydM5i/Q9RdLaKGm/BMXaSy2l9S');

SET FOREIGN_KEY_CHECKS = 1;
