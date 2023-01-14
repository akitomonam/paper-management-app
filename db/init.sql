-- CREATE DATABASE IF NOT EXISTS vgs CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
-- CREATE USER IF NOT EXISTS 'user'@'%' IDENTIFIED BY 'root';
-- GRANT ALL PRIVILEGES ON vgs.* TO 'user'@'%';
-- SET session wait_timeout=259200;
-- SET global max_allowed_packet=10485760;

-- FLUSH PRIVILEGES;

-- CREATE TABLE IF NOT EXISTS vgs.hoge1 (id int auto_increment NOT NULL PRIMARY KEY, name varchar(10), index(id));
DROP DATABASE IF EXISTS test_database;
CREATE DATABASE test_database;
USE test_database;

CREATE TABLE users (
id INT NOT NULL AUTO_INCREMENT,
username VARCHAR(255) NOT NULL,
password VARCHAR(255) NOT NULL,
file_path VARCHAR(255) NULL,
PRIMARY KEY (id));

CREATE TABLE papers (
id INT NOT NULL AUTO_INCREMENT,
title VARCHAR(255) NULL,
author VARCHAR(255) NULL,
publisher VARCHAR(255) NULL,
year INT NULL,
file_name VARCHAR(255) NOT NULL,
file_path VARCHAR(255) NOT NULL,
abstract TEXT NULL,
user_id INT NOT NULL,
created_at DATETIME NOT NULL,
PRIMARY KEY (id));


CREATE TABLE favorites (
id INT NOT NULL AUTO_INCREMENT,
paper_id INT NOT NULL,
user_id INT NOT NULL,
rating INT NOT NULL,
PRIMARY KEY (id));

CREATE TABLE keywords (
id INT NOT NULL AUTO_INCREMENT,
paper_id INT NOT NULL,
keyword VARCHAR(255) NOT NULL,
PRIMARY KEY (id));

CREATE TABLE comments (
id INT NOT NULL AUTO_INCREMENT,
paper_id INT NOT NULL,
user_id INT NOT NULL,
content TEXT NOT NULL,
created_at DATETIME NOT NULL,
PRIMARY KEY (id));

CREATE TABLE sessions (
session_token CHAR(128) NOT NULL,
user_id INT NOT NULL,
PRIMARY KEY (session_token));
