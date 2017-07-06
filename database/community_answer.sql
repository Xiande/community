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
-- Table structure for table `answer`
--

DROP TABLE IF EXISTS `answer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `answer` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `QuestionId` int(11) DEFAULT NULL,
  `UserName` varchar(45) DEFAULT NULL,
  `DisplayName` varchar(45) DEFAULT NULL,
  `IsBest` tinyint(1) DEFAULT '0',
  `BestTime` datetime DEFAULT NULL,
  `IsExpert` tinyint(1) DEFAULT '0',
  `ExpertTime` datetime DEFAULT NULL,
  `VotedCount` int(11) DEFAULT '0',
  `CreateDate` datetime DEFAULT CURRENT_TIMESTAMP,
  `CreateBy` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `answer`
--

LOCK TABLES `answer` WRITE;
/*!40000 ALTER TABLE `answer` DISABLE KEYS */;
INSERT INTO `answer` VALUES (1,21,'admin','Administrator',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-02 05:45:17','admin'),(2,21,'admin','Administrator',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-02 07:17:08','admin'),(3,21,'guest2','Guest2',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-03 08:14:29','guest2'),(4,23,'guest2','Guest2',1,'2016-11-09 06:18:03',0,'0000-00-00 00:00:00',2,'2016-11-06 02:21:22','guest2'),(6,23,'admin','Administrator',0,'0000-00-00 00:00:00',1,'2016-11-09 06:07:27',1,'2016-11-09 04:53:00','admin'),(7,23,'admin','Administrator',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',1,'2016-11-09 04:53:05','admin'),(8,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:23:47','guxiande'),(9,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:23:52','guxiande'),(10,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:23:55','guxiande'),(11,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:23:58','guxiande'),(12,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:24:02','guxiande'),(13,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:24:05','guxiande'),(14,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:24:10','guxiande'),(15,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:24:16','guxiande'),(16,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:24:21','guxiande'),(17,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:24:24','guxiande'),(18,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:24:27','guxiande'),(19,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',0,'2016-11-09 06:24:29','guxiande'),(20,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',1,'2016-11-09 06:24:34','guxiande'),(21,23,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',1,'2016-11-09 06:24:43','guxiande'),(22,22,'guxiande','Guxiande',0,'0000-00-00 00:00:00',0,'0000-00-00 00:00:00',1,'2016-11-11 05:06:45','guxiande');
/*!40000 ALTER TABLE `answer` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-17 15:41:23
