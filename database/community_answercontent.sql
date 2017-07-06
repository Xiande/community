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
-- Table structure for table `answercontent`
--

DROP TABLE IF EXISTS `answercontent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `answercontent` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `AnswerId` varchar(45) DEFAULT NULL,
  `AnswerContent` text,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `answercontent`
--

LOCK TABLES `answercontent` WRITE;
/*!40000 ALTER TABLE `answercontent` DISABLE KEYS */;
INSERT INTO `answercontent` VALUES (1,'1','<p>aaa</p>\n'),(2,'2','<p>bbb</p>\n'),(3,'3','<p>dfdf</p>\n'),(4,'4','<p>真的不能加啊</p>\n'),(6,'6','<p>a</p>\n'),(7,'7','<p>b</p>\n'),(8,'8','<p>df</p>\n'),(9,'9','<p>s</p>\n'),(10,'10','<p>s</p>\n'),(11,'11','<p>d</p>\n'),(12,'12','<p>f</p>\n'),(13,'13','<p>g</p>\n'),(14,'14','<p>h</p>\n'),(15,'15','<p>i</p>\n'),(16,'16','<p>1</p>\n'),(17,'17','<p>2</p>\n'),(18,'18','<p>3</p>\n'),(19,'19','<p>4</p>\n'),(20,'20','<p>5</p>\n'),(21,'21','<p>6</p>\n'),(22,'22','<p>dd</p>\n');
/*!40000 ALTER TABLE `answercontent` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-17 15:41:20
