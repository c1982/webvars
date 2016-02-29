-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.6.20 - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             9.3.0.4984
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping database structure for ripeips
CREATE DATABASE IF NOT EXISTS `ripeips` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `ripeips`;


-- Dumping structure for table ripeips.company
CREATE TABLE IF NOT EXISTS `company` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Netname` varchar(50) NOT NULL,
  `Status` varchar(50) DEFAULT NULL,
  `Descr` varchar(512) DEFAULT NULL,
  `Country` varchar(10) DEFAULT NULL,
  `Organization` varchar(512) DEFAULT NULL,
  `Origin` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Netname` (`Netname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Data exporting was unselected.


-- Dumping structure for table ripeips.iplist
CREATE TABLE IF NOT EXISTS `iplist` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `cid` int(11) NOT NULL,
  `ipaddr` varchar(50) NOT NULL,
  `isUp` bit(1) DEFAULT NULL,
  `isWindows` bit(1) DEFAULT NULL,
  `isLinux` bit(1) DEFAULT NULL,
  `hasWebServer` bit(1) DEFAULT NULL,
  `isPlesk` bit(1) DEFAULT NULL,
  `isCpanel` bit(1) DEFAULT NULL,
  `isWSP` bit(1) DEFAULT NULL,
  `isDirectAdmin` bit(1) DEFAULT NULL,
  `ServerHeader` varchar(50) DEFAULT NULL,
  `PoweredByHeader` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `ipaddr` (`ipaddr`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Data exporting was unselected.
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
