CREATE DATABASE IF NOT EXISTS AuthServer;

USE AuthServer;

CREATE TABLE login_details)
	Username varchar(100) NOT NULL,
	PRIMARY KEY(Username),
	Password binary(100) NOT NULL
);