CREATE DATABASE IF NOT EXISTS requestdb;

USE requestdb;

CREATE TABLE IF NOT EXISTS requests (
  username VARCHAR(50) NOT NULL PRIMARY KEY,
  password VARCHAR(100) NOT NULL
);