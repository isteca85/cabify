CREATE USER IF NOT EXISTS web IDENTIFIED BY 'cabify-challenge';
grant all privileges on *.* to web with grant option;

CREATE DATABASE IF NOT EXISTS cabify;

CREATE TABLE cabify.cars (
        ID INT NOT NULL,
        seats INT NOT NULL,
        ID_Journey INT NULL,
        PRIMARY KEY (ID)
);      

CREATE TABLE cabify.journey (
        ID INT NOT NULL,
        people INT NOT NULL,
        status TINYINT NOT NULL,
        time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (ID)
);      

