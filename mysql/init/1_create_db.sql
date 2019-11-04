DROP DATABASE IF EXISTS yuri_bot;
CREATE DATABASE IF NOT EXISTS yuri_bot;
USE yuri_bot;

CREATE TABLE IF NOT EXISTS feed (id int(11) PRIMARY KEY AUTO_INCREMENT,title varchar(191) not null unique,link varchar(191) not null unique,published varchar(191) not null,description varchar(191) not null)