CREATE DATABASE IF NOT EXISTS mydatabase;
USE mydatabase;

CREATE TABLE IF NOT EXISTS Requests (
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(100),
  description VARCHAR(300),
  status INT,
);