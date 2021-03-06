-- MySQL Script generated by MySQL Workbench
-- Сб 04 июн 2022 17:44:13
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema wb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema wb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `wb` ;
USE `wb` ;

-- -----------------------------------------------------
-- Table `wb`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `wb`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `expires` TIMESTAMP NOT NULL,
  `type` VARCHAR(4) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `wb`.`verifieds`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `wb`.`verifieds` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `token` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_verified_users1_idx` (`user_id` ASC) VISIBLE,
  UNIQUE INDEX `idverified_UNIQUE` (`id` ASC) VISIBLE,
  CONSTRAINT `fk_verified_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `wb`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `wb`.`resets`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `wb`.`resets` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `token` VARCHAR(100) NULL,
  `user_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_pass_users1_idx` (`user_id` ASC) VISIBLE,
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  CONSTRAINT `fk_pass_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `wb`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `wb`.`actions`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `wb`.`actions` (
  `idactions` INT NOT NULL AUTO_INCREMENT,
  `action` VARCHAR(45) NULL,
  `users_id` INT NOT NULL,
  PRIMARY KEY (`idactions`),
  INDEX `fk_actions_users1_idx` (`users_id` ASC) VISIBLE,
  CONSTRAINT `fk_actions_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `wb`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
