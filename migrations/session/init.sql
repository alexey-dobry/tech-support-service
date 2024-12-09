CREATE DATABASE IF NOT EXISTS sessiondb;

USE sessiondb;

CREATE TABLE IF NOT EXISTS sessions (
  manager_id INT NOT NULL,
  is_free BOOLEAN NOT NULL,
  client_id INT NOT NULL
);