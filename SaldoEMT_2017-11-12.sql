# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: localhost (MySQL 5.7.20)
# Database: SaldoEMT
# Generation Time: 2017-11-12 16:15:27 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table BusLineType
# ------------------------------------------------------------

DROP TABLE IF EXISTS `BusLineType`;

CREATE TABLE `BusLineType` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `imageUrl` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `BusLineType` WRITE;
/*!40000 ALTER TABLE `BusLineType` DISABLE KEYS */;

INSERT INTO `BusLineType` (`id`, `imageUrl`)
VALUES
	(1,'https://s3.eu-central-1.amazonaws.com/saldo-emt/Zona+Urbana.png'),
	(2,'https://s3.eu-central-1.amazonaws.com/saldo-emt/Aeroport.png'),
	(3,'https://s3.eu-central-1.amazonaws.com/saldo-emt/Port.png');

/*!40000 ALTER TABLE `BusLineType` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table BusLineTypeFare
# ------------------------------------------------------------

DROP TABLE IF EXISTS `BusLineTypeFare`;

CREATE TABLE `BusLineTypeFare` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `busLineTypeId` int(11) unsigned NOT NULL,
  `fareId` int(11) unsigned NOT NULL,
  `cost` float NOT NULL,
  PRIMARY KEY (`id`),
  KEY `busLineTypeId` (`busLineTypeId`),
  KEY `fareId` (`fareId`),
  CONSTRAINT `buslinetypefare_ibfk_1` FOREIGN KEY (`busLineTypeId`) REFERENCES `BusLineType` (`id`),
  CONSTRAINT `buslinetypefare_ibfk_2` FOREIGN KEY (`fareId`) REFERENCES `Fare` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `BusLineTypeFare` WRITE;
/*!40000 ALTER TABLE `BusLineTypeFare` DISABLE KEYS */;

INSERT INTO `BusLineTypeFare` (`id`, `busLineTypeId`, `fareId`, `cost`)
VALUES
	(1,1,1,0.8),
	(2,1,2,1.15),
	(3,1,3,0.45),
	(4,1,4,0.3),
	(5,1,5,0),
	(6,1,6,0.3),
	(7,1,7,0.3),
	(8,1,8,0),
	(9,1,9,0.3),
	(10,1,10,13),
	(11,1,11,29),
	(18,2,1,1),
	(19,2,2,5),
	(20,2,3,0.67),
	(21,2,4,0.67),
	(22,2,5,0.67),
	(23,2,6,0.67),
	(24,2,7,0.67),
	(25,2,8,0.67),
	(26,2,9,0.67),
	(28,3,1,0.8),
	(29,3,2,1.15),
	(30,3,3,0.45),
	(31,3,4,0.3),
	(32,3,5,0),
	(33,3,6,0.3),
	(34,3,7,0.3),
	(35,3,8,0),
	(36,3,9,0.3);

/*!40000 ALTER TABLE `BusLineTypeFare` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table BusLineTypeTranslation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `BusLineTypeTranslation`;

CREATE TABLE `BusLineTypeTranslation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `busLineTypeId` int(11) unsigned NOT NULL,
  `language` int(11) unsigned NOT NULL,
  `name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `busLineTypeId` (`busLineTypeId`),
  KEY `language` (`language`),
  CONSTRAINT `buslinetypetranslation_ibfk_1` FOREIGN KEY (`busLineTypeId`) REFERENCES `BusLineType` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `buslinetypetranslation_ibfk_2` FOREIGN KEY (`language`) REFERENCES `Language` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `BusLineTypeTranslation` WRITE;
/*!40000 ALTER TABLE `BusLineTypeTranslation` DISABLE KEYS */;

INSERT INTO `BusLineTypeTranslation` (`id`, `busLineTypeId`, `language`, `name`)
VALUES
	(1,1,1,'Zona Urbana'),
	(2,2,1,'Aeropuerto'),
	(3,3,1,'Puerto');

/*!40000 ALTER TABLE `BusLineTypeTranslation` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Fare
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Fare`;

CREATE TABLE `Fare` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `days` int(2) unsigned DEFAULT NULL,
  `rides` int(2) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `Fare` WRITE;
/*!40000 ALTER TABLE `Fare` DISABLE KEYS */;

INSERT INTO `Fare` (`id`, `days`, `rides`)
VALUES
	(1,NULL,NULL),
	(2,NULL,NULL),
	(3,NULL,NULL),
	(4,NULL,NULL),
	(5,NULL,NULL),
	(6,NULL,NULL),
	(7,NULL,NULL),
	(8,NULL,NULL),
	(9,NULL,NULL),
	(10,30,20),
	(11,30,50);

/*!40000 ALTER TABLE `Fare` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table FareNameTranslation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `FareNameTranslation`;

CREATE TABLE `FareNameTranslation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `fareId` int(11) unsigned NOT NULL,
  `language` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `fareId` (`fareId`,`language`),
  KEY `language` (`language`),
  CONSTRAINT `farenametranslation_ibfk_1` FOREIGN KEY (`fareId`) REFERENCES `Fare` (`id`),
  CONSTRAINT `farenametranslation_ibfk_2` FOREIGN KEY (`language`) REFERENCES `Language` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `FareNameTranslation` WRITE;
/*!40000 ALTER TABLE `FareNameTranslation` DISABLE KEYS */;

INSERT INTO `FareNameTranslation` (`id`, `name`, `fareId`, `language`)
VALUES
	(1,'Residente',1,1),
	(2,'No residente',2,1),
	(3,'Estudiante y universitario',3,1),
	(4,'Carnet Verde',4,1),
	(5,'Carnet Gran A',5,1),
	(6,'Carnet Gran B',6,1),
	(7,'Familia Numerosa',7,1),
	(8,'Menores entre 5 y 12 años',8,1),
	(9,'Menores entre 13 y 16 años',9,1),
	(10,'Abono 20 Viajes/30 días Residente',10,1),
	(11,'Abono 50 viajes/30 días residente',11,1);

/*!40000 ALTER TABLE `FareNameTranslation` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table Language
# ------------------------------------------------------------

DROP TABLE IF EXISTS `Language`;

CREATE TABLE `Language` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(2) NOT NULL DEFAULT '',
  `name` varchar(11) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `Language` WRITE;
/*!40000 ALTER TABLE `Language` DISABLE KEYS */;

INSERT INTO `Language` (`id`, `code`, `name`)
VALUES
	(1,'es','spanish'),
	(2,'ca','catalan'),
	(3,'en','english'),
	(4,'de','german'),
	(5,'fr','french');

/*!40000 ALTER TABLE `Language` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
