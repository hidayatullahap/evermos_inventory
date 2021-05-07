-- Valentina Studio --
-- MySQL dump --
-- ---------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
-- ---------------------------------------------------------


-- CREATE DATABASE "evermos" -------------------------------
CREATE DATABASE IF NOT EXISTS `evermos` CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `evermos`;
-- ---------------------------------------------------------


-- CREATE TABLE "inventories" ----------------------------------
CREATE TABLE `inventories`( 
	`id` Int( 11 ) AUTO_INCREMENT NOT NULL,
	`product_id` Int( 255 ) NOT NULL,
	`quantity` Int( 255 ) NOT NULL,
	PRIMARY KEY ( `id` ) )
CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB
AUTO_INCREMENT = 1;
-- -------------------------------------------------------------


-- Dump data of "inventories" ------------------------------
BEGIN;

INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '1', '1', '9' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '2', '2', '0' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '3', '22', '10' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '4', '81', '5' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '5', '14', '5' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '6', '169', '5' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '7', '230', '5' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '8', '393', '5' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '9', '35', '0' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '10', '476', '0' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '11', '423', '0' );
INSERT INTO `inventories`(`id`,`product_id`,`quantity`) VALUES ( '12', '461', '-45' );
COMMIT;
-- ---------------------------------------------------------


-- CREATE INDEX "index_product_id" -----------------------------
CREATE INDEX `index_product_id` USING BTREE ON `inventories`( `product_id` );
-- -------------------------------------------------------------


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
-- ---------------------------------------------------------


