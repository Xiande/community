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
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Username` varchar(100) NOT NULL,
  `Email` varchar(100) DEFAULT NULL,
  `PhotoUrl` varchar(100) DEFAULT NULL,
  `Website` varchar(256) DEFAULT NULL,
  `Location` varchar(256) DEFAULT NULL,
  `Tagline` varchar(256) DEFAULT NULL,
  `Bio` varchar(256) DEFAULT NULL,
  `Twitter` varchar(256) DEFAULT NULL,
  `Weibo` varchar(256) DEFAULT NULL,
  `GitHubUsername` varchar(256) DEFAULT NULL,
  `JoinedAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `Follow` varchar(5000) DEFAULT NULL,
  `Fans` varchar(5000) DEFAULT NULL,
  `IsSuperuser` tinyint(1) DEFAULT '0',
  `IsActive` tinyint(1) DEFAULT '0',
  `ValidateCode` varchar(45) DEFAULT NULL,
  `ResetCode` varchar(45) DEFAULT NULL,
  `UIndex` int(11) DEFAULT NULL,
  `Password` varchar(100) NOT NULL,
  `DisplayName` varchar(45) NOT NULL,
  `IsModerator` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=354 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (349,'admin','','default_boy.png','','','','','','','','0000-00-00 00:00:00','','',1,1,'','',0,'e10adc3949ba59abbe56e057f20f883e','Administrator',0),(350,'guest2','','a9b2d9faa64411e6bf4b9890969c5384.png','','','','','','','','0000-00-00 00:00:00','','',0,1,'','',0,'e10adc3949ba59abbe56e057f20f883e','guest2',0),(351,'guxiande','','c6eeea4aa64411e6bf4b9890969c5384.png','','','','','','','','0000-00-00 00:00:00','','',0,1,'','',0,'e10adc3949ba59abbe56e057f20f883e','Guxiande',0),(352,'guest1','','','','','','','','','','0000-00-00 00:00:00','','',0,1,'','',0,'e10adc3949ba59abbe56e057f20f883e','guest1',0),(353,'ddd','123456@pw.com','','','','','','','','','0000-00-00 00:00:00','','',0,1,'','',0,'e10adc3949ba59abbe56e057f20f883e','ddd',0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-17 15:41:19
