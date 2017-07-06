CREATE DATABASE  IF NOT EXISTS `community` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `community`;
-- MySQL dump 10.13  Distrib 5.7.12, for Win64 (x86_64)
--
-- Host: localhost    Database: community
-- ------------------------------------------------------
-- Server version	5.7.16-enterprise-commercial-advanced-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `question`
--

DROP TABLE IF EXISTS `question`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `question` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserName` varchar(45) DEFAULT NULL,
  `DisplayName` varchar(45) DEFAULT NULL,
  `Title` varchar(200) DEFAULT NULL,
  `Tags` varchar(100) DEFAULT NULL,
  `EmailNotice` tinyint(1) DEFAULT NULL,
  `ViewedCount` int(11) DEFAULT NULL,
  `BestCount` int(11) DEFAULT NULL,
  `AnswersCount` int(11) DEFAULT NULL,
  `VotedCount` int(11) DEFAULT NULL,
  `IsExpert` tinyint(1) DEFAULT NULL,
  `ExpertTime` datetime DEFAULT NULL,
  `LanguageType` varchar(20) DEFAULT NULL,
  `CreateDate` datetime DEFAULT NULL,
  `CreateBy` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `question`
--

LOCK TABLES `question` WRITE;
/*!40000 ALTER TABLE `question` DISABLE KEYS */;
INSERT INTO `question` VALUES (21,'admin','Administrator','t4','SharePoint',1,25,0,1,0,0,'0000-00-00 00:00:00','en','2016-11-02 04:37:18','admin'),(22,'admin','Administrator','发的辅导费空间里的饭卡的','GoLang,testtag',1,11,0,1,1,0,'0000-00-00 00:00:00','cn','2016-11-02 07:42:58','admin'),(23,'guest2','Guest2','不能加表情','常用',1,85,0,17,6,1,'2016-11-09 06:07:27','cn','2016-11-06 02:20:41','guest2'),(27,'guxiande','Guxiande','test tag','SharePoint,GoLang,Go,noq',1,4,0,0,0,0,'0000-00-00 00:00:00','en','2016-11-11 05:06:16','guxiande');
/*!40000 ALTER TABLE `question` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-17 15:41:21
