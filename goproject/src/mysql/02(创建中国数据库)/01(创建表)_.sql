
-- 创建数据库
create database china charset=utf8;

-- 创建省表
CREATE TABLE T_Province(
  ProID INT PRIMARY KEY AUTO_INCREMENT,
  ProName VARCHAR(50) NOT NULL,
  ProSort INT,
  ProRemark VARCHAR(50)
);




-- 创建市表
CREATE TABLE T_City
(
    CityID INT Primary KEY AUTO_INCREMENT,
    CityName VARCHAR(50)  NOT NULL,
    ProID INT,
    CitySort INT
);


-- 创建县表
CREATE TABLE  T_District
(
    Id INT PRIMARY KEY AUTO_INCREMENT,
    DisName VARCHAR(30) NOT NULL,
    CityID INT NOT NULL,
    DisSort INT
);
