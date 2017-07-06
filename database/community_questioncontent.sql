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
-- Table structure for table `questioncontent`
--

DROP TABLE IF EXISTS `questioncontent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `questioncontent` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `QuestionId` int(11) DEFAULT NULL,
  `QuestionContent` text,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `questioncontent`
--

LOCK TABLES `questioncontent` WRITE;
/*!40000 ALTER TABLE `questioncontent` DISABLE KEYS */;
INSERT INTO `questioncontent` VALUES (1,12,'<p>test</p>\n'),(4,18,'<p>tt</p>\n'),(5,19,'<p>t2</p>\n'),(6,20,'<p>t3</p>\n'),(7,21,'<p>t4</p>\n'),(8,22,'<p>大幅度发<img alt=\"\" src=\"/uploadimages/e4a2f812-a0cf-11e6-808b-9890969c5384.png\" style=\"height:438px; width:780px\" /></p>\n'),(9,23,'<p>不能加表情</p>\n'),(10,24,'<p>del</p>\n'),(12,27,'<p>teate</p>\n');
/*!40000 ALTER TABLE `questioncontent` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-17 15:41:24
