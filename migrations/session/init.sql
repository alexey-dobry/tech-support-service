CREATE DATABASE IF NOT EXISTS sessiondb;

USE sessiondb;

CREATE TABLE IF NOT EXISTS sessions (
  managerid INT NOT NULL,
  clientid INT NOT NULL
);