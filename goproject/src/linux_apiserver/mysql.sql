


CREATE DATABASE IF not EXISTS db_apiserver;

use db_apiserver;


CREATE TABLE tb_users(
   id INTEGER PRIMARY KEY AUTO_INCREMENT,
   username VARCHAR(200) NOT NULL,
   password VARCHAR(200) NOT NULL,
   createdAt timestamp null default null,
   updatedAt timestamp null default null,
   deletedAt timestamp null default null,
   UNIQUE(username)
);



-- 下面的这个是测试使用的

CREATE TABLE users(
   id INTEGER PRIMARY KEY AUTO_INCREMENT,
   name VARCHAR(200) NOT NULL,
   password VARCHAR(200) NOT NULL,
   createdAt timestamp null default null,
   updatedAt timestamp null default null,
   deletedAt timestamp null default null,
   UNIQUE(name)
);
