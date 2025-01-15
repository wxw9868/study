-- MySQL dump 10.13  Distrib 9.0.1, for macos14 (arm64)
--
-- Host: localhost    Database: part_time
-- ------------------------------------------------------
-- Server version	9.0.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `pt_ad`
--

DROP TABLE IF EXISTS `pt_ad`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_ad` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ad_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '广告名称',
  `ad_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '链接地址',
  `img_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '图片地址',
  `start_time` datetime(3) NOT NULL COMMENT '开始时间·',
  `end_time` datetime(3) NOT NULL COMMENT '结束时间',
  `is_show` tinyint NOT NULL DEFAULT '1' COMMENT '是否显示: 1显示 2不显示',
  `is_free` tinyint NOT NULL COMMENT '是否免费: 1是 2不是',
  `sort` bigint NOT NULL DEFAULT '0' COMMENT '排序',
  `ad_type` tinyint NOT NULL COMMENT '广告类型: 1banner广告 2普通广告',
  PRIMARY KEY (`id`),
  KEY `idx_pt_ad_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_ad`
--

LOCK TABLES `pt_ad` WRITE;
/*!40000 ALTER TABLE `pt_ad` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_ad` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_article`
--

DROP TABLE IF EXISTS `pt_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_article` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章标题',
  `author` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章作者',
  `content` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章内容',
  PRIMARY KEY (`id`),
  KEY `idx_pt_article_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_article`
--

LOCK TABLES `pt_article` WRITE;
/*!40000 ALTER TABLE `pt_article` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_article` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_coupon`
--

DROP TABLE IF EXISTS `pt_coupon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_coupon` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '优惠券名称',
  `level` smallint DEFAULT NULL COMMENT '优惠券等级',
  `full_amount` decimal(10,2) NOT NULL COMMENT '金额',
  `send_amount` decimal(10,2) NOT NULL COMMENT '满多少金额送多少: 例如满100送10',
  `is_use` tinyint NOT NULL DEFAULT '1' COMMENT '是否使用: 1已使用 2没使用',
  `desc` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '优惠券介绍',
  PRIMARY KEY (`id`),
  KEY `idx_pt_coupon_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_coupon`
--

LOCK TABLES `pt_coupon` WRITE;
/*!40000 ALTER TABLE `pt_coupon` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_coupon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_coupon_log`
--

DROP TABLE IF EXISTS `pt_coupon_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_coupon_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `order_id` bigint unsigned NOT NULL COMMENT '订单ID',
  `coupon_id` bigint unsigned NOT NULL COMMENT '优惠券ID',
  PRIMARY KEY (`id`),
  KEY `idx_pt_coupon_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_coupon_log`
--

LOCK TABLES `pt_coupon_log` WRITE;
/*!40000 ALTER TABLE `pt_coupon_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_coupon_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_education`
--

DROP TABLE IF EXISTS `pt_education`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_education` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sort` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_pt_education_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_education`
--

LOCK TABLES `pt_education` WRITE;
/*!40000 ALTER TABLE `pt_education` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_education` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_job`
--

DROP TABLE IF EXISTS `pt_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_job` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `cat_id` bigint unsigned NOT NULL COMMENT '职位分类',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '职位名称',
  `number` bigint unsigned NOT NULL COMMENT '招聘人数',
  `is_discuss` tinyint NOT NULL DEFAULT '1' COMMENT '是否面议: 1否 2是',
  `max_salary` double NOT NULL COMMENT '最高薪资',
  `min_salary` double NOT NULL COMMENT '最低薪资',
  `settlement_id` bigint unsigned NOT NULL COMMENT '薪资结算方式ID',
  `start_time` datetime(3) NOT NULL COMMENT '开始工作时间',
  `end_time` datetime(3) NOT NULL COMMENT '结束工作时间',
  `description` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '职位描述',
  `province` int NOT NULL COMMENT '工作所在地：省',
  `city` int NOT NULL COMMENT '工作所在地：市',
  `district` int NOT NULL COMMENT '工作所在地：区/县',
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '工作详细地址',
  `liaison` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '联系人',
  `mobile` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号',
  `wechat_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '微信号',
  `status` tinyint NOT NULL COMMENT '审核状态: 1待审核 2审核通过 3审核失败',
  `comment` text COLLATE utf8mb4_unicode_ci COMMENT '备注',
  `is_show` tinyint NOT NULL DEFAULT '2' COMMENT '是否上线: 1上线 2下线',
  `hide_time` datetime(3) DEFAULT NULL COMMENT '下线时间',
  `show_time` datetime(3) DEFAULT NULL COMMENT '上线时间',
  `is_top` tinyint NOT NULL DEFAULT '2' COMMENT '是否置顶: 1置顶 2取消',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_pt_job_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_job`
--

LOCK TABLES `pt_job` WRITE;
/*!40000 ALTER TABLE `pt_job` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_job` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_job_category`
--

DROP TABLE IF EXISTS `pt_job_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_job_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '职位分类名称',
  `icon` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '职位分类图标',
  `image` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '分类图片路径',
  `sort` bigint NOT NULL COMMENT '排序',
  `is_show` tinyint NOT NULL DEFAULT '0' COMMENT '是否显示: 2不显示 1显示',
  `recommend` tinyint NOT NULL DEFAULT '0' COMMENT '金刚位显示: 2不显示 1显示',
  PRIMARY KEY (`id`),
  KEY `idx_pt_job_category_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_job_category`
--

LOCK TABLES `pt_job_category` WRITE;
/*!40000 ALTER TABLE `pt_job_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_job_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_job_deliver`
--

DROP TABLE IF EXISTS `pt_job_deliver`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_job_deliver` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `job_id` bigint unsigned NOT NULL COMMENT '岗位ID',
  `hire_user_id` bigint unsigned NOT NULL COMMENT '发布岗位的用户ID',
  `apply_user_id` bigint unsigned NOT NULL COMMENT '申请岗位的用户ID',
  `status` tinyint NOT NULL COMMENT '状态: 1被查看 2已录取 3已拒绝 4待处理',
  PRIMARY KEY (`id`),
  KEY `idx_pt_job_deliver_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_job_deliver`
--

LOCK TABLES `pt_job_deliver` WRITE;
/*!40000 ALTER TABLE `pt_job_deliver` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_job_deliver` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_job_delivery_meter`
--

DROP TABLE IF EXISTS `pt_job_delivery_meter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_job_delivery_meter` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `delivery_id` bigint unsigned NOT NULL COMMENT '用户投递表ID',
  `pay_status` tinyint NOT NULL COMMENT '扣费状态: 1已扣费 2没有扣费',
  PRIMARY KEY (`id`),
  KEY `idx_pt_job_delivery_meter_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_job_delivery_meter`
--

LOCK TABLES `pt_job_delivery_meter` WRITE;
/*!40000 ALTER TABLE `pt_job_delivery_meter` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_job_delivery_meter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_job_promotion`
--

DROP TABLE IF EXISTS `pt_job_promotion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_job_promotion` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `job_id` bigint unsigned NOT NULL COMMENT '岗位ID',
  `promotion_fee` decimal(10,2) NOT NULL COMMENT '推广费用',
  `promotion_type` tinyint NOT NULL COMMENT '推广类型: 1banner广告 2普通广告 3列表置顶',
  `promotion_status` tinyint NOT NULL COMMENT '申请状态: 1待审核 2通过 3不通过',
  `start_time` datetime NOT NULL COMMENT '开始时间',
  `end_time` datetime NOT NULL COMMENT '结束时间',
  PRIMARY KEY (`id`),
  KEY `idx_pt_job_promotion_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_job_promotion`
--

LOCK TABLES `pt_job_promotion` WRITE;
/*!40000 ALTER TABLE `pt_job_promotion` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_job_promotion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_job_settlement_type`
--

DROP TABLE IF EXISTS `pt_job_settlement_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_job_settlement_type` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `settlement_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '结算名称',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '说明',
  `sort` int NOT NULL DEFAULT '1' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_pt_job_settlement_type_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_job_settlement_type`
--

LOCK TABLES `pt_job_settlement_type` WRITE;
/*!40000 ALTER TABLE `pt_job_settlement_type` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_job_settlement_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_order`
--

DROP TABLE IF EXISTS `pt_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `order_sn` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '订单号',
  `order_status` tinyint NOT NULL COMMENT '订单状态: 1已确认 2已取消 3已完成 4已作废',
  `job_promotion_id` bigint unsigned DEFAULT '0' COMMENT '职位推广ID',
  `order_amount` decimal(10,2) NOT NULL COMMENT '订单金额',
  `order_type` tinyint NOT NULL COMMENT '订单类型: 1充值 2消费 3推广 4提现',
  `order_desc` longtext COLLATE utf8mb4_unicode_ci COMMENT '订单说明',
  `payment_method` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '支付方式: alipay支付宝支付 wechatpay微信支付',
  `payment_status` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态: 1已支付 2待支付 3支付失败',
  `payment_time` datetime(3) NOT NULL COMMENT '支付时间',
  `transaction_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方平台交易流水号',
  PRIMARY KEY (`id`),
  KEY `idx_pt_order_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_order`
--

LOCK TABLES `pt_order` WRITE;
/*!40000 ALTER TABLE `pt_order` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_region`
--

DROP TABLE IF EXISTS `pt_region`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_region` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `level` tinyint DEFAULT '0',
  `parent_id` int DEFAULT NULL,
  `region_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_pt_region_deleted_at` (`deleted_at`),
  KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_region`
--

LOCK TABLES `pt_region` WRITE;
/*!40000 ALTER TABLE `pt_region` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_region` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_settings`
--

DROP TABLE IF EXISTS `pt_settings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_settings` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `wechat_show` tinyint NOT NULL DEFAULT '2' COMMENT '微信显示: 1展示 2隐藏',
  `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '投递一次岗位的收费价格',
  `customer_service` longtext COLLATE utf8mb4_unicode_ci COMMENT '客户服务',
  PRIMARY KEY (`id`),
  KEY `idx_pt_settings_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_settings`
--

LOCK TABLES `pt_settings` WRITE;
/*!40000 ALTER TABLE `pt_settings` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_settings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_university`
--

DROP TABLE IF EXISTS `pt_university`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_university` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `school_name` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '学校名称',
  `school_identifier` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '学校标识码',
  `competent_department` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '主管部门',
  `location` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '所在地',
  `note` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_pt_university_school_name` (`school_name`),
  UNIQUE KEY `idx_pt_university_school_identifier` (`school_identifier`),
  KEY `idx_pt_university_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_university`
--

LOCK TABLES `pt_university` WRITE;
/*!40000 ALTER TABLE `pt_university` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_university` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_user`
--

DROP TABLE IF EXISTS `pt_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `mobile` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
  `user_type` tinyint NOT NULL DEFAULT '0' COMMENT '用户类型: 1普通用户 2企业用户 3管理员',
  `usertype` tinyint NOT NULL DEFAULT '0' COMMENT '用户类型: 1普通用户 2企业用户 3管理员',
  `sex` tinyint NOT NULL DEFAULT '0' COMMENT '性别: 0保密 1男 2女',
  `wechat_id` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '微信号',
  `qq` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'QQ',
  `education` tinyint DEFAULT NULL COMMENT '学历: 1小学 2初中 3高中 4大专 5本科 6研究生',
  `degree` tinyint DEFAULT NULL COMMENT '学历: 1学士 2硕士 3博士',
  `intro` longtext COLLATE utf8mb4_unicode_ci COMMENT '简介',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_pt_user_username` (`username`),
  UNIQUE KEY `idx_pt_user_mobile` (`mobile`),
  KEY `idx_pt_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_user`
--

LOCK TABLES `pt_user` WRITE;
/*!40000 ALTER TABLE `pt_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_user_balance`
--

DROP TABLE IF EXISTS `pt_user_balance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_user_balance` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `give_money` decimal(10,2) DEFAULT '0.00' COMMENT '满赠的额度',
  PRIMARY KEY (`id`),
  KEY `idx_pt_user_balance_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_user_balance`
--

LOCK TABLES `pt_user_balance` WRITE;
/*!40000 ALTER TABLE `pt_user_balance` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_user_balance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_user_balance_log`
--

DROP TABLE IF EXISTS `pt_user_balance_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_user_balance_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `money` decimal(10,2) NOT NULL COMMENT '金额',
  `payment_method` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '支付方式: alipay支付宝支付 wechatpay微信支付',
  `action` tinyint NOT NULL COMMENT '行为: 1充值 2消费 3推广 4提现 5岗位扣费',
  `amount` decimal(10,2) NOT NULL COMMENT '金额',
  `payment_status` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态: 1已支付 2待支付 3支付失败',
  `payment_time` datetime(3) NOT NULL COMMENT '支付时间',
  `transaction_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方平台交易流水号',
  PRIMARY KEY (`id`),
  KEY `idx_pt_user_balance_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_user_balance_log`
--

LOCK TABLES `pt_user_balance_log` WRITE;
/*!40000 ALTER TABLE `pt_user_balance_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_user_balance_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_user_certification`
--

DROP TABLE IF EXISTS `pt_user_certification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_user_certification` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `type` tinyint DEFAULT NULL COMMENT '认证类型: 1企业用户 2个人用户',
  `industry` bigint unsigned DEFAULT NULL COMMENT '所属行业: job_category表ID',
  `company_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '公司名称',
  `company_logo` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '公司logo',
  `business_license` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '营业执照',
  `realname` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '真实姓名',
  `idcard` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '身份证号',
  `idcard_front` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '身份证正面',
  `idcard_back` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '身份证反面',
  `province` bigint unsigned DEFAULT NULL COMMENT '省',
  `city` bigint unsigned DEFAULT NULL COMMENT '市',
  `district` bigint unsigned DEFAULT NULL COMMENT '区/县',
  `address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '详细地址',
  `intro` text COLLATE utf8mb4_unicode_ci COMMENT '简介',
  `status` tinyint DEFAULT '1' COMMENT '状态: 1待审核 2审核通过 3审核失败',
  `note` text COLLATE utf8mb4_unicode_ci COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_pt_user_certification_company_name` (`company_name`),
  UNIQUE KEY `idx_pt_user_certification_idcard` (`idcard`),
  KEY `idx_pt_user_certification_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_user_certification`
--

LOCK TABLES `pt_user_certification` WRITE;
/*!40000 ALTER TABLE `pt_user_certification` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_user_certification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_user_login_log`
--

DROP TABLE IF EXISTS `pt_user_login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_user_login_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `info_id` bigint DEFAULT NULL COMMENT '访问ID',
  `login_name` longtext COLLATE utf8mb4_unicode_ci COMMENT '登录账号',
  `ipaddr` longtext COLLATE utf8mb4_unicode_ci COMMENT '登录IP地址',
  `login_location` longtext COLLATE utf8mb4_unicode_ci COMMENT '登录地点',
  `browser` longtext COLLATE utf8mb4_unicode_ci COMMENT '浏览器类型',
  `os` longtext COLLATE utf8mb4_unicode_ci COMMENT '操作系统',
  `status` bigint DEFAULT NULL COMMENT '登录状态(1成功 2失败)',
  `msg` longtext COLLATE utf8mb4_unicode_ci COMMENT '提示消息',
  `login_time` datetime(3) DEFAULT NULL COMMENT '登录时间',
  `module` longtext COLLATE utf8mb4_unicode_ci COMMENT '登录模块',
  PRIMARY KEY (`id`),
  KEY `idx_pt_user_login_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_user_login_log`
--

LOCK TABLES `pt_user_login_log` WRITE;
/*!40000 ALTER TABLE `pt_user_login_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_user_login_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pt_userinfo`
--

DROP TABLE IF EXISTS `pt_userinfo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pt_userinfo` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `realname` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '真实姓名',
  `sex` tinyint NOT NULL DEFAULT '0' COMMENT '性别: 0保密 1男 2女',
  `region` int NOT NULL COMMENT '区',
  `education_id` bigint unsigned NOT NULL COMMENT '最高学历',
  `wechat_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '微信号',
  `is_student` tinyint DEFAULT NULL COMMENT '是否学生: 1是 2不是',
  `is_import` tinyint DEFAULT '0' COMMENT '是否导入: 1是 0不是',
  PRIMARY KEY (`id`),
  KEY `idx_pt_userinfo_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pt_userinfo`
--

LOCK TABLES `pt_userinfo` WRITE;
/*!40000 ALTER TABLE `pt_userinfo` DISABLE KEYS */;
/*!40000 ALTER TABLE `pt_userinfo` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'part_time'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-01-15 10:48:50
