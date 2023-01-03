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
PRIMARY KEY (id),
FOREIGN KEY (user_id) REFERENCES users (id));


CREATE TABLE favorites (
id INT NOT NULL AUTO_INCREMENT,
paper_id INT NOT NULL,
user_id INT NOT NULL,
PRIMARY KEY (id),
FOREIGN KEY (paper_id) REFERENCES papers (id),
FOREIGN KEY (user_id) REFERENCES users (id));

CREATE TABLE keywords (
id INT NOT NULL AUTO_INCREMENT,
paper_id INT NOT NULL,
keyword VARCHAR(255) NOT NULL,
PRIMARY KEY (id),
FOREIGN KEY (paper_id) REFERENCES papers (id));

CREATE TABLE sessions (
session_token CHAR(128) NOT NULL,
user_id INT NOT NULL,
PRIMARY KEY (session_token),
FOREIGN KEY (user_id) REFERENCES users (id));
