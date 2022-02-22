-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        8.0.21 - MySQL Community Server - GPL
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  11.2.0.6213
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- 导出  表 aah-form-based-auth.permissions 结构
CREATE TABLE IF NOT EXISTS `permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_permissions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  aah-form-based-auth.permissions 的数据：~2 rows (大约)
DELETE FROM `permissions`;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
INSERT INTO `permissions` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`) VALUES
	(1, '2022-02-18 14:32:22.585', '2022-02-18 14:32:22.585', NULL, 'users:manage:view'),
	(2, '2022-02-18 14:32:22.585', '2022-02-18 14:32:22.585', NULL, 'users:*');
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;

-- 导出  表 aah-form-based-auth.roles 结构
CREATE TABLE IF NOT EXISTS `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `quota` float DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  aah-form-based-auth.roles 的数据：~3 rows (大约)
DELETE FROM `roles`;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `quota`) VALUES
	(1, '2022-02-18 14:32:22.521', '2022-02-18 14:32:22.521', NULL, 'user', 0),
	(2, '2022-02-18 14:32:22.521', '2022-02-18 14:32:22.521', NULL, 'manager', 0),
	(3, '2022-02-18 14:32:22.521', '2022-02-18 14:32:22.521', NULL, 'administrator', 0);
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;

-- 导出  表 aah-form-based-auth.users 结构
CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `first_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `last_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` char(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `is_locked` tinyint(1) DEFAULT NULL,
  `is_expried` tinyint(1) DEFAULT NULL,
  `created_at` datetime(6) DEFAULT NULL,
  `updated_at` datetime(6) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  aah-form-based-auth.users 的数据：~4 rows (大约)
DELETE FROM `users`;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `first_name`, `last_name`, `email`, `password`, `is_locked`, `is_expried`, `created_at`, `updated_at`, `deleted_at`) VALUES
	('225ab8c1-86f6-48f8-ac01-24dbf66cdd00', 'West', 'West Corner', 'user2@aahframework.org', '$2a$10$skXZnY3mGAFChPGSXOdrLuGaI/IYYD7cGNLhKenHTSgfrANoZDjrW', 1, 0, '2022-02-18 14:32:22.972000', '2022-02-18 14:32:22.972000', NULL),
	('43c47c56-787b-4e97-8031-4c524b3c2d33', 'South', 'South Corner', 'user3@aahframework.org', '$2a$10$SeY0Jflx/xhLtfRE/21VC.9kVHqCWDkswgWA/w8IY4fskmZRe6FPi', 0, 0, '2022-02-18 14:32:23.089000', '2022-02-18 14:32:23.510000', NULL),
	('8cf75dcc-a06f-41de-acf2-0c074ee03202', 'East', 'East Corner', 'user1@aahframework.org', '$2a$10$ctDLJmTUL4AClzSHfD3sHeQVd1O0zxVR8BjlxawkT6O5694O8sknW', 0, 0, '2022-02-18 14:32:22.863000', '2022-02-18 14:32:23.369000', NULL),
	('d4272bf3-01de-401d-9079-c81caef92f39', 'Admin', 'Admin Corner', 'admin@aahframework.org', '$2a$10$q/TPSfyuuYgQ5j0U8IqcQuvJLmZ9.4M3TiW5JF8uI2vJLt2CRFBNq', 0, 0, '2022-02-18 14:32:22.750000', '2022-02-18 14:32:23.792000', NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

-- 导出  表 aah-form-based-auth.users_permissions 结构
CREATE TABLE IF NOT EXISTS `users_permissions` (
  `user_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `permission_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`permission_id`),
  KEY `fk_users_permissions_permission` (`permission_id`),
  CONSTRAINT `fk_users_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`),
  CONSTRAINT `fk_users_permissions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  aah-form-based-auth.users_permissions 的数据：~2 rows (大约)
DELETE FROM `users_permissions`;
/*!40000 ALTER TABLE `users_permissions` DISABLE KEYS */;
INSERT INTO `users_permissions` (`user_id`, `permission_id`) VALUES
	('8cf75dcc-a06f-41de-acf2-0c074ee03202', 1),
	('d4272bf3-01de-401d-9079-c81caef92f39', 2);
/*!40000 ALTER TABLE `users_permissions` ENABLE KEYS */;

-- 导出  表 aah-form-based-auth.users_roles 结构
CREATE TABLE IF NOT EXISTS `users_roles` (
  `user_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`),
  KEY `fk_users_roles_role` (`role_id`),
  CONSTRAINT `fk_users_roles_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `fk_users_roles_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 正在导出表  aah-form-based-auth.users_roles 的数据：~5 rows (大约)
DELETE FROM `users_roles`;
/*!40000 ALTER TABLE `users_roles` DISABLE KEYS */;
INSERT INTO `users_roles` (`user_id`, `role_id`) VALUES
	('43c47c56-787b-4e97-8031-4c524b3c2d33', 1),
	('8cf75dcc-a06f-41de-acf2-0c074ee03202', 1),
	('d4272bf3-01de-401d-9079-c81caef92f39', 1),
	('8cf75dcc-a06f-41de-acf2-0c074ee03202', 2),
	('d4272bf3-01de-401d-9079-c81caef92f39', 3);
/*!40000 ALTER TABLE `users_roles` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
