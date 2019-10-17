


CREATE DATABASE IF EXISTS db_apiserver;

use db_apiserver;


CREATE TABLE tb_users(
   id INTEGER PRIMARY KEY AUTO_INCREMENT,
   username VARCHAR(200) NOT NULL,
   PASSWORD VARCHAR(200) NOT NULL,
   CreatedAt timestamp null default null,
   UpdatedAt timestamp null default null,
   DeletedAt timestamp null default null,
   UNIQUE(username)
);

