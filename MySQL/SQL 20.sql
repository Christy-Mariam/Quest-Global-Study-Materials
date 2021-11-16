/*SQL Joints*/

CREATE DATABASE db2;

USE db2;

CREATE TABLE country(
country_code VARCHAR(10) NOT NULL,
country_name VARCHAR(100),
PRIMARY KEY (country_code)
);

INSERT INTO country
VALUES
('IN', 'India'),
('SL', 'Sri Lanka'),
('PK', 'Pakistan'),
('BN', 'Bangladesh');

SELECT * FROM country;

CREATE TABLE capital(
capital_id VARCHAR(10) NOT NULL,
country_code VARCHAR(10),
capital_name VARCHAR(100),
PRIMARY KEY (capital_id),
FOREIGN KEY (country_code) REFERENCES country(country_code)
);

INSERT INTO capital
VALUES
('c1', 'IN', 'New Delhi'),
('c2','PK', 'Islamabad');

SELECT * FROM capital;

