

-- 创建数据库 mydb
create database mydb charset=utf8

-- 创建person数据库
create table person(
    id int primary key auto_increment,
    name varchar(100),
    age int ,
    rmb float ,
    gender bool,
    birthday date
);