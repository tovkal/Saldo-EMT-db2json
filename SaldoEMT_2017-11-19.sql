# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: localhost (MySQL 5.7.20)
# Database: SaldoEMT
# Generation Time: 2017-11-19 21:03:37 +0000
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
  `displayBusLineTypeName` bit(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `BusLineType` WRITE;
/*!40000 ALTER TABLE `BusLineType` DISABLE KEYS */;

INSERT INTO `BusLineType` (`id`, `imageUrl`, `displayBusLineTypeName`)
VALUES
	(1,'https://s3.eu-central-1.amazonaws.com/saldo-emt/Zona+Urbana.png',b'0'),
	(2,'https://s3.eu-central-1.amazonaws.com/saldo-emt/Aeroport.png',b'1'),
	(3,'https://s3.eu-central-1.amazonaws.com/saldo-emt/Port.png',b'1');

/*!40000 ALTER TABLE `BusLineType` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
